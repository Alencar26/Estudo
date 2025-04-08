# GIN-Api - API REST com Go e Gin

Este é um projeto de API REST desenvolvido em Go utilizando o framework Gin e GORM para gerenciamento de alunos.

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)

## 📋 Pré-requisitos

- Go instalado
- PostgreSQL instalado e configurado
- Variáveis de ambiente configuradas (arquivo .env)

## 🔧 Configuração do Ambiente

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
```

2. Configure as variáveis de ambiente no arquivo `.env`:
```env
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_DB=root
PGADMIN_DEFAULT_EMAIL=email@golang.com
PGADMIN_DEFAULT_PASSWORD=123456
DB_HOST=localhost
DB_USERNAME=root
DB_PASSWORD=root
DB_DATABASE=root
DB_PORT=5432
```

## 🏃‍♂️ Executando o Projeto

1. Instale as dependências:
```bash
go mod tidy
```

2. Execute o projeto:
```bash
go run main.go
```

O servidor será iniciado na porta 5000.

## 📌 Endpoints da API

### Alunos

#### GET `/alunos`
- Retorna todos os alunos cadastrados

#### GET `/alunos/:id`
- Retorna um aluno específico pelo ID

#### POST `/alunos`
- Cria um novo aluno
- Payload de exemplo:
```json
{
    "nome": "Nome do Aluno",
    "cpf": "123.456.789-00",
    "rg": "12.345.678-9"
}
```

#### DELETE `/alunos/:id`
- Remove um aluno pelo ID

### Outros Endpoints

#### GET `/:nome`
- Retorna uma saudação com o nome fornecido

## 📁 Estrutura do Projeto

```
GIN-Api/
├── database/
│   └── db.go
├── handles/
│   └── handles.go
├── models/
│   └── aluno.go
├── routes/
│   └── routes.go
├── .env
├── main.go
└── README.md
```

## 🛠️ Funcionalidades

- CRUD completo de alunos
- Conexão com banco de dados PostgreSQL
- Migrations automáticas com GORM
- Tratamento de erros
- Respostas HTTP padronizadas

## ⚙️ Modelo de Dados

### Aluno
- ID (automático)
- Nome
- CPF
- RG
- Timestamps (CreatedAt, UpdatedAt, DeletedAt)

## 👥 Autor

André Alencar

## 📄 Licença

Este projeto está sob a licença MIT - consulte o arquivo LICENSE.md para detalhes.
