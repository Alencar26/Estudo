# Migrations & SQLC

### Migrations

Instalar o [golang migrate](https://github.com/golang-migrate/migrate)

Para  iniciar uma migration no projeto devemos usar o comando:
```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```
Será criado dois arquivos *up* e *down*, o arquivo *up* é usado na hora da criação da migration e o arquivo *down* é utilizado na destruição da migration.
Insira código SQL nos arquivos.

Após inserir os SQL nos arquivos, suba uma banco de dados e execute o seguinte comando:
```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
```
Para desfazer a migration:
```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```
### SQLC

O SqlC faz um mapeamento de todas as querys que serão utilizadas no projeto e gera automaticamente o código GO para lidar com elas

Instalação [slqc github](https://github.com/sqlc-dev/sqlc?tab=readme-ov-file)

Na raiz do projeto crie um arquivo `sqlc.yaml`. 

Exemplo do conteúdo do arquivo:
```yaml
version: "2"
sql:
- schema: "sql/migrations"
  queries: "sql/queries"
  engine: "mysql"
  gen:
    go:
      package: "db"
      out: "internal/db"
```
Para gerar (dentro da raiz do projeto) execute no terminal:
```bash
sqlc generate
```
