package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
  ID int `gorm:"primaryKey"`
  Name string
  Price float64
  gorm.Model
}

//gorm.Model = Adiciona atributos de gerenciamento ao nosso modelo.

func main() {
  dns := "root:root@tcp(localhost:3306)/godb" 
  db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&Product{}) //GORM cria a tabela automaticamente.

  //Insert into db ------------------------
  db.Create(&Product{
    Name: "Laptop",
    Price: 1000.00,
  })

  //Insert in batch -----------------------
  products := []Product{
    {Name: "P1", Price: 2.99},
    {Name: "P2", Price: 23.50},
    {Name: "P3", Price: 7.99},
  }

  db.Create(&products)

  //Select one -----------------------------
  var p1 Product
  db.First(&p1, 1) //return ID 1
  fmt.Println(p1)

  //Select by name -------------------------
  var p2 Product
  db.First(&p2, "name = ?", "P2")
  fmt.Println(p2)

  //Select All ----------------------------
  var allProducts []Product
  db.Find(&allProducts)
  for _, p := range allProducts {
    fmt.Println(p)
  }

  //Select All with limit ----------------
  // Offset = pagination
  // Limit = return n results
  db.Limit(3).Offset(2).Find(&allProducts)
  fmt.Println(allProducts)


  //Using Where ------------------------
  db.Where("price < ?", 20.00).Find(&allProducts)
  fmt.Println(allProducts)

  //Using LIKE --------------------------
  db.Where("name LIKE ?", "%3%").Find(&allProducts)
  fmt.Println(allProducts)

  //Update -----------------------------
  var p Product
  db.First(&p, 1)
  p.Name = "Novo Valor"
  db.Save(&p)

  //Delete ------------------------------
  //db.Delete(&p)
}
