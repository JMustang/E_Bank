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

---

### Usando SQLC

O SQLC é uma ferramenta poderosa que converte consultas SQL em código Go fortemente tipado. Isso facilita o desenvolvimento de aplicações que utilizam bancos de dados relacionais, evitando erros de tipagem e melhorando a integração entre o banco de dados e o código Go.

#### 1. Configuração do SQLC

Primeiro, você precisa configurar o SQLC no seu projeto. Crie um arquivo `sqlc.yaml` na raiz do projeto com a seguinte configuração:

```yaml
version: "1"
packages:
  - name: "db"
    path: "db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_interface: false
```

Neste exemplo:

- `name`: Nome do pacote gerado.
- `path`: Caminho onde o código gerado será salvo.
- `queries`: Diretório onde suas consultas SQL serão armazenadas.
- `schema`: Diretório onde seus arquivos de esquema SQL (migrações) estão localizados.
- `engine`: Tipo de banco de dados usado (neste caso, PostgreSQL).

#### 2. Escrevendo Consultas SQL

Crie um diretório chamado `db/query/` e adicione suas consultas SQL lá. Por exemplo, para criar uma consulta que busca uma conta bancária pelo ID, crie um arquivo chamado `account.sql`:

```sql
-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE id = $1 LIMIT 1;
```

Nesta consulta:

- `-- name: GetAccount :one` define o nome da função Go gerada (`GetAccount`) e especifica que a consulta retornará uma única linha (`:one`).

#### 3. Gerando Código Go

Com as consultas SQL escritas, você pode gerar o código Go executando o seguinte comando na raiz do projeto:

```bash
sqlc generate
```

Este comando criará um pacote Go no caminho especificado no arquivo `sqlc.yaml` (`db/sqlc/`). O pacote conterá funções para executar suas consultas SQL diretamente a partir do código Go.

#### 4. Usando o Código Gerado

Depois de gerar o código, você pode usar as funções em seu código Go. Por exemplo, para buscar uma conta pelo ID:

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
    "github.com/usuario/e_bank/db/sqlc"
)

func main() {
    conn, err := sql.Open("postgres", "postgresql://usuario:senha@localhost:5432/e_bank_db?sslmode=disable")
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    store := db.NewStore(conn)

    accountID := int64(1)
    account, err := store.GetAccount(context.Background(), accountID)
    if err != nil {
        log.Fatal("cannot get account:", err)
    }

    fmt.Printf("Account: %+v\n", account)
}
```

Neste exemplo:

- `db.NewStore(conn)` cria uma nova instância da estrutura que contém as funções SQL geradas.
- `store.GetAccount(context.Background(), accountID)` executa a consulta SQL `GetAccount` gerada pelo SQLC.

### 5. Atualizando Consultas SQL

Se precisar adicionar ou modificar consultas SQL, basta editar os arquivos no diretório `db/query/` e executar novamente o comando `sqlc generate` para regenerar o código Go atualizado.

---
