package mocks

import (
	"testing"
  "errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
  mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
  args := m.Called(tax)
  return args.Error(0)
}

func TestCalculateTaxAndSave(t *testing.T) {
  repository := &TaxRepositoryMock{}
  repository.On("SaveTax", 10.0).Return(nil).Once()
  repository.On("SaveTax", 0.0).Return(errors.New("Error Saving Tax"))
  
  err := CalculateTaxAndSave(1000.0, repository)
  assert.Nil(t, err)

  err = CalculateTaxAndSave(0.0, repository)
  assert.Error(t, err, "Erro Saving Tax")

  repository.AssertExpectations(t)
  repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
