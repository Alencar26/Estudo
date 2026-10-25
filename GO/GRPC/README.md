# GRPC

[Documentoação Oficial do gRPC para GO](https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc).

Introdução ao protocolo gRPC utilizando HTTP 2.

gRPC faz o uso de protocol buffers, [saiba mais - doc](https://protobuf.dev/).

Se faz necessário criar um arquivo com a extensão `.proto` para definies os schemas das nossas entidades.
Após criação desse arquivo deve-se rodar o comando abaixo:

```bash
protoc --go_out=. --go-grpc_out=. proto/entities.proto
```

o parâmetro `--go_out` vai gerar os arquivos das entidades e messages definidos no arquivo `.proto` em um diretório definido.
o parâmetro `--go-grpc_out` vai gerar arquivos e intefaces para trabalhar com gRPC em um diretório definido.
Passamos também o caminho do nosso arquivo `.proto`

**ATENÇÃO!! Sempre que houver modificações no arquivo .proto deve ser executado o comando acima.**

### Client gRPC

O client que vamos usar é o **Evans**. (Client escrito em Go ara lidar com requisições gRPC)

[Evans Doc](https://github.com/ktr0731/evans)

Comando para rodar o client:

```bash
evans -r repl
```
No Evans é preciso sempre iniciar escolhendo o pacote que será utilizado com o comando `package nome_pacote`
Depois é preciso informar qual o serviço com o comando `service nome_serviço`
E por fim a escolha das chamadas com o comando `call nome_da_chamada`.

Em uma chamada de Stream n Evans para finalizar precione `CTRL+d`
