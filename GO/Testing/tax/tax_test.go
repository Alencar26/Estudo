package tax

import (
	"testing"
)

//COMANDOS
//  go test | Só executa os testes sem detalhes
//  go test -v | Execução dos testes com detalhes das funções testadas
//  go test -coverprofile=coverage.out | Executa os testes e mostra a cobertura de teste no código
//  go tool cover -html=coverage.out | Vai exportar para HTML o código que está coberto por testes e mostrar o que não está coberto.

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
    {0.0, 0.0},
  }

  for _, item := range table {
    result := CalculateTax(item.amount)
    if result != item.expcted {
     t.Errorf("Expected %f but got %f", item.expcted, result)
    }
  }
}
