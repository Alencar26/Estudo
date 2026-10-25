package main

import "fmt"

func makeIdGenerator() func() int {
  id := 0
  return func() int { //essa função interna é o clusure
    id++
    return id
  }
}

func main()  {
  generateID := makeIdGenerator()
  generateID2 := makeIdGenerator() // nesse caso aqui não é compartilhado com id da instância acima.

  fmt.Println(
    generateID(), //mesmo chando várias vezes - Ele vai compartilhar o estado e valor da vairável id
    generateID(), //mesmo que a vairável id não esteja no escopo da closure ela utiliza-se da memória para manipular o valor
    generateID(),
    )

  fmt.Println(
    generateID2(),
    generateID2(),
    )
}
