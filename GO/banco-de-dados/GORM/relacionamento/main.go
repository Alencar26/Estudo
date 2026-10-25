package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
  ID        int `gorm:"promaryKey"`
  Name      string
  Products  []Product
}

type Product struct {
  ID            int `gorm:"primaryKey"`
  Name          string
  Price         float64
  CategoryID    int
  SerialNumber  SerialNumber
  Category      Category
  gorm.Model
}

type SerialNumber struct {
  ID        int `gorm:"primaryKey"`
  Number    string
  ProductID int
}

//gorm.Model = Adiciona atributos de gerenciamento ao nosso modelo.

func main() {
  dns := "root:root@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local" 
  db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}) //GORM cria a tabela automaticamente.

  //CREATE ---------------------------------

  //create category
  category := Category{Name: "Eletronicos"}
  db.Create(&category)

  //create product
  db.Create(&Product{
    Name: "Mouse",
    Price: 100,
    CategoryID: category.ID,
  })

  db.Create(&SerialNumber{
    Number: "202506091752",
    ProductID: 1, //colocando valor na mão pois é só um exemplo
  })

  // FIND -------------------------------
  var products []Product
  db.Preload("Category").Preload("SerialNumber").Find(&products)
  for _, p := range products {
    fmt.Printf("Produto: %v | SerialNumber: %v | Categoria: %v\n", p.Name, p.SerialNumber.Number, p.Category.Name)
  }

  var categories []Category
  err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
  if err != nil {
    panic(err)
  }
  for _, c := range categories {
    fmt.Println(c.Name,"-------------")
    for _, p := range c.Products {
      fmt.Println(p.Name)
      fmt.Println(p.SerialNumber.Number)
    }  
  }
}
