package main

import "fmt"

func main() {
  fmt.Println(calculate(2,3))
}

/*

REGRAS DE CALCULO DA NOTAÇÃO BIG (O)

Atribuição (x = 5) = 1 unidade de tempo;
Retornos (return x) = 1 unidade de tempo;
Operações Aritiméticas (x + y) = 1 unidade de tempo;
Operações lógicas (x && y) = 1 unidade de tempo;
Pequenas operações = 1 unidade de tempo;


*/


func calculate(x, y int) int {
  result := x + y //nessa linha temos 1 + 1 + 1 + 1 = 4
  return result // 1 + 1 = 2 
}

//Função acima é Big O(1)

// A função acima vai gastar sempre ≃6 unidades de tempo para executar;
// Significando que ela é constante. Não importa a entrada, sempre será 6 unidades de tempo.


//Na função abaixo eu estou obtendo um valor específico do array sem percorrer ele.
func getValueByKey(array []int, k int) int {
  return array[k] // 1 + 1 + 1 = ≃ T3 (Constante) = Big O(1) 
}
