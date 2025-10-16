# AWS CLI

### Estrutura dos comandos
```Bash
#aws [options] <command> <subcommand> [parameters]
aws [opção extras] <recurso da aws> <ação desse resurso> [parametros]

#0 = APP CLI
#1 = Options
#2 = Command
#3 = Subcommand
#4 = Parameters

#exemplo
aws --endpoint-url=http://localhost:4566 s3  ls  --profile localstack
[0] [                    1             ] [2] [3] [        4          ]
```
---

### Configuração de usuário na AWS

```bash
aws configure --profile <username>
```

É possível verificar a configuração no diretório `~/.aws/`.
Nesse diretório deveconer dois arquivos `config` e `credentials`.

As configurações de usuário realizadas pello comando acima podem ser alteradas diretamente nos respectivos arquivos.

---

### Comando na AWS com localstack

Para usar a AWS com localstack é necessário informar o endpoint do localstack. 
E informar o user do localstak

```bash
aws --endpoint http://localhost:4566 --profile localstack s3 mb s3://meubuckets3 
```

