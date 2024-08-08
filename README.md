# E_Bank Backend

Database diagram
![db diagram](Simple_Bank.png)

---

Este é o backend para o E_Bank, uma aplicação bancária que permite a gestão de contas, transferências e consultas de saldo. O backend é desenvolvido em Go e utiliza PostgreSQL como banco de dados relacional. O projeto foi configurado para ser facilmente desenvolvido e testado em diferentes ambientes, com o uso de Docker, Migrate para migrações de banco de dados, SQLC para geração de queries, Testify para testes e Dotenv para o gerenciamento de variáveis de ambiente.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal usada para desenvolver o backend.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar informações.
- **SQLC**: Ferramenta para gerar código Go a partir de consultas SQL.
- **Testify**: Framework para testes em Go.
- **Migrate**: Ferramenta para gerenciamento e migração de esquemas de banco de dados.
- **Docker**: Usado para containerizar a aplicação e facilitar o ambiente de desenvolvimento e produção.
- **Dotenv**: Para gerenciar variáveis de ambiente.

## Requisitos

Antes de começar, certifique-se de ter os seguintes softwares instalados em seu ambiente de desenvolvimento:

- [Golang](https://golang.org/dl/) (versão 1.18+)
- [PostgreSQL](https://www.postgresql.org/download/) (versão 13+)
- [Docker](https://www.docker.com/get-started)
- [Git](https://git-scm.com/downloads)

## Instalação e Configuração

### 1. Clone o Repositório

```bash
git clone https://github.com/JMustang/E_Bank.git
cd E_Bank
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
DB_SOURCE=postgresql://usuario:senha@localhost:5432/e_bank_db?sslmode=disable
DB_DRIVER=postgres
```

### 3. Instalação das Dependências

Para instalar todas as dependências necessárias, execute:

```bash
go mod tidy
```

### 4. Instalar SQLC

SQLC é usado para gerar código Go a partir de consultas SQL:

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### 5. Instalar Migrate

Migrate é usado para gerenciar migrações do banco de dados:

```bash
brew install golang-migrate   # macOS
```

ou

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 6. Instalar Testify

Testify é um framework de testes:

```bash
go get github.com/stretchr/testify
```

### 7. Executar Migrações

Para aplicar as migrações no banco de dados, use o Migrate:

```bash
migrate -path db/migration -database "$DB_SOURCE" -verbose up
```

### 8. Executar a Aplicação

Para iniciar a aplicação, utilize o comando:

```bash
go run main.go
```

### 9. Utilizar Docker

Se preferir utilizar Docker para executar a aplicação:

```bash
docker-compose up --build
```

## Testes

Para executar os testes da aplicação, utilize o comando:

```bash
go test -v ./...
```

## Contribuição

Se desejar contribuir com o projeto, faça um fork do repositório, crie um branch com suas alterações, e depois envie um pull request.

---
