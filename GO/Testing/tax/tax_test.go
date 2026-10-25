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

//COMANDO PARA RODAR O BENCHMARK
//  go test -bench=. | Rodar tests + Benchmark
//  go test -bench=. -run=^# | REGEX que vai impedir que rodem os testes juntos. Só vai rodar o benchmark
//  go test -bench=. -run=^# -count=10 | Vai rodar o benchmark 10x para verificar a média
//  go test -bench=. -run=^# -benchmem | Informações de alocação de memória
func BenchmarkCalculateTax(b *testing.B) {
  for i := 0; i < b.N; i ++ {
    CalculateTax(500.0)
  }
}

//comparação do desempenho da funcção abaixo com a função acima;
func BenchmarkCalculateTax2(b *testing.B) {
  for i := 0; i < b.N; i ++ {
    CalculateTax2(500.0)
  }
}

//FUZZING - TESTES COM VALORES ALEATÓRIOS 
//COMANDO: 
//  go test -fuzz=.
//  go test -funzz=. -fuzztime=5s -run=^#
func FuzzCalculateTax(f *testing.F) {
  //"Treinando" - passando valores de exemplo par o fuzzing
  seed := []float64{-1, -2, -2.5, 500.0, 0, 1000.0, 1001.4,999.67}
  for _, amount := range seed {
    f.Add(amount)
  }

  //Execução de fato
  f.Fuzz(func(t *testing.T, amount float64) {

    result := CalculateTax(amount)
    if amount <= 0 && result != 0 {
      t.Errorf("Expected 0 but got %f - Amount: %f", result, amount)
    }
    if amount > 20000 && result != 20.0 {
      t.Errorf("Expected 20 but got %f - Amount: %f", result, amount)
    }

  })
} 
