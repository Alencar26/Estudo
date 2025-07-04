package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	//user:passwor@tcp(127.0.0.1:3306)/nomeDoBancoDeDados

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Balinha", 2.99)
	err = insertProduct(db, *product)
	if err != nil {
		panic(err)
	}

	product.Price = 9.99
	err = updateProduct(db, *product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, "51b08b53-65d9-4272-95c7-c19b214ddd1d")
	if err != nil {
		panic(err)
	}
	fmt.Printf("O produto: %v, tem o preço de: R$ %.2f", p.Name, p.Price)

  products, err := selectAllProducts(db)
  if err != nil {
    panic(err)
  }
  for _, p := range products {
    fmt.Println(p)
  }

  err = deleteProduct(db, "0f8d13d2-e61c-4e30-8eb8-f120496871d7")
  if err != nil {
    panic(err)
  }
}

func insertProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error){
  rows, err := db.Query("select id, name, price from products;")
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var products []Product
  for rows.Next() {
    var p Product
    err = rows.Scan(&p.ID, &p.Name, &p.Price)
    if err != nil {
      return nil, err
    }
    products = append(products, p)
  }
  return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
  stmt, err := db.Prepare("delete from products where id = ?;")
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(id)
  if err != nil {
    return err
  }
  return nil
}
