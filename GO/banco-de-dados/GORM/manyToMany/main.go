package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
  ID        int `gorm:"promaryKey"`
  Name      string
  Products  []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
  ID            int `gorm:"primaryKey"`
  Name          string
  Price         float64
  Categories    []Category `gorm:"many2many:products_categories;"`
  gorm.Model
}

//gorm.Model = Adiciona atributos de gerenciamento ao nosso modelo.

func main() {
  dns := "root:root@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local" 
  db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{}, &Category{}) //GORM cria a tabela automaticamente.

  //CREATE ---------------------------------

  //create category
  category := Category{Name: "Eletronicos"}
  db.Create(&category)

  category2 := Category{Name: "Sala"}
  db.Create(&category2)

  //create product
  db.Create(&Product{
    Name: "Televisão",
    Price: 100,
    Categories: []Category{category, category2},
  })


  // FIND -------------------------------
  var products []Product
  db.Preload("Categories").Find(&products)
  for _, p := range products {
    fmt.Printf("Produto: %v | Categorias: %v\n", p.Name, p.Categories)
  }
}
