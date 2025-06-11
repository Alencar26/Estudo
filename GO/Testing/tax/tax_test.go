package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T)  {
  amount := 500.0
  expcted := 5.0

  result := CalculateTax(amount)
  
  if result != expcted {
    t.Errorf("Expected %f but got %f", expcted, result)
  }
}
