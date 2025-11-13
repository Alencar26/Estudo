# Migrations

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
