package database

import (
	"math/rand/v2"
	"testing"

	"github.com/Alencar26/Estudo/GO/APIS/internal/entity"
	"github.com/Alencar26/Estudo/GO/APIS/pkg/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createProductForTest() (*entity.Product, error) {
	product, err := entity.NewProduct("Product"+test.RandomLetter(4), rand.IntN(100))
	if err != nil {
		return nil, err
	}
	return product, nil
}

func TestCreateProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productTest, err := createProductForTest()
	assert.Nil(t, err)
	assert.NotNil(t, productTest)

	productDB := NewProduct(db)
	err = productDB.Create(productTest)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", productTest.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, productTest.Name, productFound.Name, "Nome do produto é diferente do esperado")
	assert.Equal(t, productTest.Price, productFound.Price, "Preço do produto é diferente do esperado")
	assert.Equal(t, productTest.ID, productFound.ID, "ID do produto é diferente do esperado")
}

func TestFindProductByID(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productTest, err := createProductForTest()
	assert.Nil(t, err)
	assert.NotNil(t, productTest)

	productDB := NewProduct(db)
	err = productDB.Create(productTest)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(string(productTest.ID.String()))
	assert.Nil(t, err, "Erro retornou diferente de NIL ao buscar produto por ID")
	assert.NotNil(t, productFound, "Produto retornou NIL na busca por ID")

	assert.Equal(t, productTest.Name, productFound.Name, "Nome do produto é diferente do esperado")
	assert.Equal(t, productTest.Price, productFound.Price, "Preço do produto é diferente do esperado")
	assert.Equal(t, productTest.ID, productFound.ID, "ID do produto é diferente do esperado")

}

func TestUpdateProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productTest, err := createProductForTest()
	assert.Nil(t, err)
	assert.NotNil(t, productTest)

	productDB := NewProduct(db)
	err = productDB.Create(productTest)
	assert.Nil(t, err)

	productFound, _ := productDB.FindByID(string(productTest.ID.String()))
	productFound.Name = "Test Update Name"
	productFound.Price = rand.IntN(300)

	err = productDB.Update(productFound)
	assert.Nil(t, err)

	productFoundUpdated, _ := productDB.FindByID(string(productTest.ID.String()))
	assert.Equal(t, productFound.Name, productFoundUpdated.Name, "Nome atualizado não é igual ao esperado")
	assert.Equal(t, productFound.Price, productFoundUpdated.Price, "Preço atualizado nãi é igual ao esperado")
	assert.Equal(t, productTest.ID, productFound.ID, "ID do produto é diferente do esperado")
}
func TestDeleteProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productTest, err := createProductForTest()
	assert.Nil(t, err)
	assert.NotNil(t, productTest)

	productDB := NewProduct(db)
	err = productDB.Create(productTest)
	assert.Nil(t, err)

	err = productDB.Delete(string(productTest.ID.String()))
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(string(productTest.ID.String()))
	assert.NotNil(t, err, "Erro deve retornar ao buscar produto deletado")
	assert.Nil(t, productFound, "Produto deve retornar NIL após deletar")
}

func TestFindAllProducts(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	// Create multiple products for testing
	product1, err := createProductForTest()
	assert.Nil(t, err)
	err = productDB.Create(product1)
	assert.Nil(t, err)

	product2, err := createProductForTest()
	assert.Nil(t, err)
	err = productDB.Create(product2)
	assert.Nil(t, err)

	product3, err := createProductForTest()
	assert.Nil(t, err)
	err = productDB.Create(product3)
	assert.Nil(t, err)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 3, len(products), "Deve retornar 3 produtos")
}

func TestFindAllProductsWithPagination(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	for i := 0; i < 5; i++ {
		product, err := createProductForTest()
		assert.Nil(t, err)
		err = productDB.Create(product)
		assert.Nil(t, err)
	}

	products, err := productDB.FindAll(1, 2, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products), "Primeira página deve retornar 2 produtos")

	products, err = productDB.FindAll(2, 2, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products), "Segunda página deve retornar 2 produtos")

	products, err = productDB.FindAll(3, 2, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products), "Terceira página deve retornar 1 produto")
}

func TestFindAllProductsWithSorting(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product1, _ := entity.NewProduct("ZProduct", 100)
	product2, _ := entity.NewProduct("AProduct", 200)
	product3, _ := entity.NewProduct("MProduct", 300)

	err = productDB.Create(product1)
	assert.Nil(t, err)
	err = productDB.Create(product2)
	assert.Nil(t, err)
	err = productDB.Create(product3)
	assert.Nil(t, err)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "ZProduct", products[0].Name, "Primeiro produto deve ser ZProduct em ordem crescente")

	products, err = productDB.FindAll(1, 10, "desc")
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "MProduct", products[0].Name, "Primeiro produto deve ser MProduct em ordem decrescente")
}

func TestFindProductByIDNotFound(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product, err := productDB.FindByID("non-existent-id")
	assert.NotNil(t, err, "Erro deve retornar ao buscar produto inexistente")
	assert.Nil(t, product, "Produto deve retornar NIL quando não encontrado")
}

func TestUpdateProductNotFound(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	productTest, err := createProductForTest()
	assert.Nil(t, err)

	err = productDB.Update(productTest)
	assert.NotNil(t, err, "Erro deve retornar ao tentar atualizar produto inexistente")
}

func TestDeleteProductNotFound(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	err = productDB.Delete("non-existent-id")
	assert.NotNil(t, err, "Erro deve retornar ao tentar deletar produto inexistente")
}
