package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
  tax,err := CalculateTax(1000.0)
  assert.Nil(t, err)
  assert.Equal(t,10.0,tax)

  tax, err = CalculateTax(0)
  assert.Error(t, err, "Valor deve ser diferente de zero")
  assert.Equal(t, 0.0, tax)
  assert.Contains(t, err.Error(), "deve ser diferente de zero")
}
