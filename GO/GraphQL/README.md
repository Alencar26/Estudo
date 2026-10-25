# Comandos utilizados no gqlgen

## Configuração inicial

[Documentação Oficial do gqlgen](https://gqlgen.com/)

Para iniciar um projeto com GraphQL no GO é necessários primeiro criar um projeto:

```bash
mkdir project-XPTO
cd project-XPTO
go mod init project-XPTO
```

Após isso, se faz necessário criar um arquivo chamado `tools.go` na raiz com os imports do gqlgen.
A criação pode ser manual ou via comando.

Comando para criar o `tools.go`:

```bash
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
```

Criação manual:

```go
//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
```

Após criação do arquivo, deve-se rodar o comando abaixo para usar gerar a estrutura bases para utilização do GraphQL no go.

```bash
go run github.com/99designs/gqlgen init
go mod tidy
```

Agora já possuímos a estrutura mínima necessária para usar o GraphQL.

Use o comando `go run server.go` no terminal para rodar o servidor.

---

## Alterando o Schema do GraphQL

Para começar a usar o GraphQL para nossa aplicação devemos fazer algumas alterações pois ele vem como um modelo de "Todo List".
Para inicar a altesção deve ser realializada a alteração do arquivo `./graph/model/models_gen.go` onde ficam os schema do GraphQL.

> Podemos termamais detalhes de como criar os Schemas na documentação oficial do GraphQL. (Link abaixo)

[Documentação Oficial - Lean GraphQL Schema](https://graphql.org/learn/schema/)

Após alteração do Schema, devemos rodar o comando abaixo para que seja gerada uma nova estrutura agora com o novo Schema.

```bash
go run github.com/99designs/gqlgen generate
go mod tidy
```

Apoś executar o comando, **haverá necessidade de realizar algumas correções em alguns arquivos**. Pois, não estamos mais usando os Schema do "Todo List" inicial e isso faz com que o Go reclame de algumas coisas.

Navegue para o arquivo `./graph/schema.resolvers.go` e no final dele conter um aviso de "WARNING" comentado. Apegue tudo do "WARNING" para baixo.

Nesse mesmo arquivo `./graph/schema.resolvers.go` foram criados várias funções para nos auxiliarem nas querys com GraphQL.
	Agora precisamos implementa-las.

**PS: TODA FEZ QUE HOUVER ALTERÇÃO NO CÓDIGO ONDE O GRAPHQL DEVE ENTENDER E SE ADEQUAR, ENTÃO, DEVE-SE RODAR NOVAMENTE O COMANDO:**

```bash
	go run github.com/99designs/gqlgen generate
````
