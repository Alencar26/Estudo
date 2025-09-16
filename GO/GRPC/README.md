# GRPC

Introdução ao protocolo gRPC utilizando HTTP 2.

Se faz necessário criar um arquivo com a extensão `.proto` para definies os schemas das nossas entidades.
Após criação desse arquivo deve-se rodar o comando abaixo:

```bash
protoc --go_out=. --go-grpc_out=. proto/entities.proto
```

o parâmetro `--go_out` vai gerar os arquivos das entidades e messages definidos no arquivo `.proto` em um diretório definido.
o parâmetro `--go-grpc_out` vai gerar arquivos e intefaces para trabalhar com gRPC em um diretório definido.
Passamos também o caminho do nosso arquivo `.proto`
