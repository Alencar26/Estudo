package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Produto1", 10)
	assert.Nil(t, err, "Error diferente de Nil")
	assert.NotNil(t, product, "Produto está retornando Nil")

	assert.Equal(t, "Produto1", product.Name, "Nome do produto é diferente do esperado")
	assert.Equal(t, 10, product.Price, "Preço do produto é diferente do esperado")
	assert.NotNil(t, product.ID)
}

func TestProductNilWhenReturnValidateError(t *testing.T) {
	product, err := NewProduct("", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
}

func TestProductWhenNameIdRequired(t *testing.T) {
	_, err := NewProduct("", 10)
	assert.Error(t, err, "Validação do nome com string vazia não retornou erro")
	assert.EqualError(t, err, ErrNameIsRequired.Error(), "Mensagem do erro é diferente do esperado")
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("ProductTesting", 0)
	assert.Error(t, err, "Validação do preço zerado não retornou erro")
	assert.EqualError(t, err, ErrPriceIsRequired.Error(), "Mensagem do erro é diferente do esperado")
}

func TestProductWhenInvalidPrice(t *testing.T) {
	_, err := NewProduct("ProductTesting", -2)
	assert.Error(t, err, "Validação do preço negativo não retornou erro")
	assert.EqualError(t, err, ErrInvalidPrice.Error(), "Mensagem do erro é diferente do esperado")
}

func TestProductWhenValidateIsNil(t *testing.T) {
	product, err := NewProduct("Product Test", 30)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
