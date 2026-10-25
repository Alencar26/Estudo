package main

import (
  "fmt"
  "packaging/math"
)

func main()  {

  m := math.Math{
    A: 10,
    B: 5,
  }

  fmt.Println(m.Add())
}

//DICA: Se eu quiser adicionar um pacote local que não está publicado no github ou na internet
//      Eu devo realizar o seguinte comando no terminals:
//      go mod edit -replace github.com/meu-pacote/funcionalidade/nome-do-pacote=../caminhoRelativo/doPacote
//      
//      O comando acima informa ao GO que o pacote em questão deve ser achado localmente a partir do caminho relativo que informamos após o "="
//
//      É preciso dar um go mod tidy no final de tudo.


//MELHOR SOLUÇÃO PARA PACOTES LOCAIS
//    Na pasta raiz rodar o comando go work init ./pacote1 ./pacote2 ../pacote3
//    Ele ira criar um arquivo chamado de go.work contendo a versão do Go + os pacotes locais para uso, sem a necessidade de uar o go.mod
//    
//IGNORAR PACORTES COM ERRO NO GO MOD TIDY: usar a flag -e
//
//Ex: go mod tidy -e
