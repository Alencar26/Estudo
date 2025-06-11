package main

import (	

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "gorm.io/gorm/clause"
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

  //LOCK OTIMISTA VS PESSIMISTA
  //LOCK = travar sua base de dados para alteração para que outras pessoas não utilizem os dados enquento estou usando.

  //Vamos fazer de forma manual o Lock Pessimista do primeiro registro de Categorias
  //Isso significa que ele vai acrescenter FOR UPDATE ao final da query para que aquele linha do banco fique travada para mim
  //Ela só será liberada após executar o update de fato.
  tx := db.Begin()
  var c Category
  err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
  if err != nil {
    panic(err)
  }

  c.Name = "Roupas"
  tx.Debug().Save(&c) //Debug é para printar no terminal o comando que o tx está executando
  tx.Commit()

  //Esse tipo de estratégia é super válida para sistemas de grande escada quando há muita concorrência 
  //Se faz necessário congelar algum registro no banco para alteração para que outro processo não use aquele registro.
}
