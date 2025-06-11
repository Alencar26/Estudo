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

func TestCalculateTaxBatch(t * testing.T) {
  type calcTax struct {
    amount, expcted float64
  }

  table := []calcTax{
    {500.0, 5.0},
    {1000.0, 10.0},
    {600.0, 5.0},
    {1300.0, 10.0},
  }

  for _, item := range table {
    result := CalculateTax(item.amount)
    if result != item.expcted {
     t.Errorf("Expected %f but got %f", item.expcted, result)
    }
  }
}
