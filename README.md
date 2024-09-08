# E_Bank Backend

Database diagram
![db diagram](Simple_Bank2.png)

---

Este é o backend para o E_Bank, uma aplicação bancária que permite a gestão de contas, transferências e consultas de saldo. O backend é desenvolvido em Go e utiliza PostgreSQL como banco de dados relacional. O projeto foi configurado para ser facilmente desenvolvido e testado em diferentes ambientes, com o uso de Docker, Migrate para migrações de banco de dados, SQLC para geração de queries, Testify para testes e Dotenv para o gerenciamento de variáveis de ambiente.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal usada para desenvolver o backend.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar informações.
- **SQLC**: Ferramenta para gerar código Go a partir de consultas SQL.
- **Testify**: Framework para testes em Go.
- **Bcrypt**: Usando para fazer Hashing de Senhas.
- **Migrate**: Ferramenta para gerenciamento e migração de esquemas de banco de dados.
- **Docker**: Usado para containerizar a aplicação e facilitar o ambiente de desenvolvimento e produção.
- **Viper**: Para gerenciar variáveis de ambiente.
- **PASETO**: Autenticação Baseada em Token (Alternativa ao JWT. Falam as más línguas que ~~melhor~~ mais simples. ).

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

### Configuração e Utilização do PostgreSQL

O PostgreSQL é um banco de dados relacional poderoso e flexível, utilizado como o sistema de gerenciamento de banco de dados (DBMS) para o backend do E_Bank. Abaixo, você encontrará instruções sobre como configurar e utilizar o PostgreSQL com a sua aplicação.

#### 1. Instalando o PostgreSQL

Se você ainda não tem o PostgreSQL instalado, siga as instruções abaixo para instalar:

**No Ubuntu/Debian:**

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

**No macOS (usando Homebrew):**

```bash
brew update
brew install postgresql
brew services start postgresql
```

**No Windows:**

Baixe e instale a partir do [site oficial do PostgreSQL](https://www.postgresql.org/download/windows/).

#### 2. Configurando o Banco de Dados

Após a instalação, siga os passos abaixo para configurar o banco de dados:

1. **Inicie o serviço PostgreSQL** (se ainda não estiver em execução):

   - No Linux: `sudo service postgresql start`
   - No macOS: `brew services start postgresql`
   - No Windows: O serviço deve iniciar automaticamente após a instalação.

2. **Acesse o prompt de comando do PostgreSQL:**

```bash
sudo -u postgres psql
```

3. **Crie um novo banco de dados para o E_Bank:**

```sql
CREATE DATABASE e_bank_db;
```

4. **Crie um usuário para o banco de dados:**

```sql
CREATE USER usuario_e_bank WITH ENCRYPTED PASSWORD 'senha_segura';
```

5. **Conceda privilégios ao usuário criado:**

```sql
GRANT ALL PRIVILEGES ON DATABASE e_bank_db TO usuario_e_bank;
```

6. **Saia do prompt de comando do PostgreSQL:**

```sql
\q
```

#### 3. Configurando a Conexão com o Banco de Dados no Projeto

Você precisará configurar a string de conexão ao banco de dados no seu arquivo `.env`. O conteúdo deve ser algo como:

```env
DB_SOURCE=postgresql://usuario_e_bank:senha_segura@localhost:5432/e_bank_db?sslmode=disable
DB_DRIVER=postgres
```

#### 4. Usando PostgreSQL com Go

No seu código Go, você pode utilizar o banco de dados configurado da seguinte forma:

1. **Instale o driver PostgreSQL para Go (`pq`):**

```bash
go get github.com/lib/pq
```

2. **Exemplo de Conexão e Consulta no Banco de Dados:**

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    dbSource = "postgresql://usuario_e_bank:senha_segura@localhost:5432/e_bank_db?sslmode=disable"
)

func main() {
    db, err := sql.Open("postgres", dbSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer db.Close()

    var id int
    var owner string
    var balance float64
    var currency string

    err = db.QueryRow("SELECT id, owner, balance, currency FROM accounts WHERE id = $1", 1).Scan(&id, &owner, &balance, &currency)
    if err != nil {
        log.Fatal("query error:", err)
    }

    fmt.Printf("ID: %d, Owner: %s, Balance: %.2f, Currency: %s\n", id, owner, balance, currency)
}
```

Nesse exemplo:

- **`sql.Open`**: Abre a conexão com o banco de dados usando a string de conexão especificada.
- **`db.QueryRow`**: Executa uma consulta SQL e mapeia o resultado para as variáveis correspondentes.

#### 5. Migrações de Banco de Dados

Para aplicar as migrações do banco de dados com Migrate, você pode usar o seguinte comando:

```bash
migrate -path db/migration -database "$DB_SOURCE" -verbose up
```

Este comando aplicará todas as migrações pendentes no seu banco de dados PostgreSQL.

---

Este exemplo fornece uma configuração completa para o PostgreSQL no contexto do seu projeto Go, incluindo a instalação, configuração, e exemplos de uso no código.

---

### Usando Viper para Gerenciamento de Configurações

[Viper](https://github.com/spf13/viper) é uma biblioteca poderosa para gerenciamento de configurações em Go. Ele suporta a leitura de variáveis de ambiente, arquivos de configuração em vários formatos (JSON, TOML, YAML, HCL, INI), e muito mais. Neste tutorial, você aprenderá como integrar o Viper ao seu projeto E_Bank para gerenciar as configurações da aplicação.

#### 1. Instalando o Viper

Primeiro, adicione o Viper ao seu projeto Go:

```bash
go get github.com/spf13/viper
```

#### 2. Configurando o Viper

No seu projeto, você pode usar o Viper para carregar variáveis de ambiente e arquivos de configuração. Abaixo está um exemplo básico de como configurar o Viper para carregar variáveis de ambiente e um arquivo `.env`:

```go
package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
)

func initConfig() {
    // Configurando o Viper para ler o arquivo .env
    viper.SetConfigFile(".env")

    // Tentando ler o arquivo de configuração
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file, %s", err)
    }

    // Configurando o Viper para ler variáveis de ambiente
    viper.AutomaticEnv()

    // Definindo valores padrão
    viper.SetDefault("DB_DRIVER", "postgres")
}

func main() {
    initConfig()

    // Acessando as configurações
    dbDriver := viper.GetString("DB_DRIVER")
    dbSource := viper.GetString("DB_SOURCE")

    fmt.Printf("Database Driver: %s\n", dbDriver)
    fmt.Printf("Database Source: %s\n", dbSource)
}
```

#### 3. Estrutura do Arquivo `.env`

Aqui está um exemplo de como o arquivo `.env` pode ser estruturado:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://usuario_e_bank:senha_segura@localhost:5432/e_bank_db?sslmode=disable
PORT=8080
```

No código acima:

- **`viper.SetConfigFile(".env")`**: Especifica o arquivo `.env` como a fonte de configuração.
- **`viper.ReadInConfig()`**: Lê o arquivo de configuração.
- **`viper.AutomaticEnv()`**: Permite que o Viper leia variáveis de ambiente diretamente.
- **`viper.SetDefault()`**: Define um valor padrão para uma configuração caso ela não seja encontrada.

#### 4. Acessando Configurações

Depois de configurar o Viper, você pode acessar as configurações em qualquer parte do seu código usando `viper.Get<Type>("key")`, onde `<Type>` é o tipo da configuração (`String`, `Int`, `Bool`, etc.) e `"key"` é o nome da chave de configuração.

Exemplo:

```go
port := viper.GetInt("PORT")
fmt.Printf("Server running on port %d\n", port)
```

#### 5. Uso Avançado do Viper

Você pode usar o Viper para carregar configurações de múltiplos arquivos, hierarquia de configurações, e mais. Abaixo está um exemplo de como carregar um arquivo YAML:

**Arquivo de Configuração YAML (`config.yaml`):**

```yaml
server:
  port: 8080
database:
  driver: postgres
  source: postgresql://usuario_e_bank:senha_segura@localhost:5432/e_bank_db?sslmode=disable
```

**Código para Carregar o YAML:**

```go
viper.SetConfigName("config")
viper.SetConfigType("yaml")
viper.AddConfigPath(".") // Diretório atual

if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("Error reading config file, %s", err)
}

port := viper.GetInt("server.port")
dbDriver := viper.GetString("database.driver")

fmt.Printf("Server running on port %d\n", port)
fmt.Printf("Database Driver: %s\n", dbDriver)
```

#### 6. Exemplo Completo

Aqui está um exemplo completo que combina leitura de arquivos `.env` e suporte a variáveis de ambiente:

```go
package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
)

func initConfig() {
    viper.SetConfigFile(".env")
    viper.SetConfigType("env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file, %s", err)
    }

    viper.SetDefault("PORT", 8080)
    viper.SetDefault("DB_DRIVER", "postgres")
}

func main() {
    initConfig()

    port := viper.GetInt("PORT")
    dbDriver := viper.GetString("DB_DRIVER")
    dbSource := viper.GetString("DB_SOURCE")

    fmt.Printf("Server running on port %d\n", port)
    fmt.Printf("Database Driver: %s\n", dbDriver)
    fmt.Printf("Database Source: %s\n", dbSource)
}
```

Este tutorial mostra como usar o Viper para gerenciar configurações de maneira eficiente no seu projeto E_Bank. Com ele, você pode facilmente ajustar configurações sem alterar o código, tornando sua aplicação mais flexível e fácil de manter.

---

<!-- ~/. -->

## Go Mock

### Usando GoMock para Mockar Dependências nos Testes

[GoMock](https://github.com/golang/mock) é uma biblioteca popular para criar mocks em testes de unidades no Go. Ela permite simular o comportamento de interfaces, ajudando a testar o código de maneira isolada e eficiente. Neste tutorial, você aprenderá como integrar e utilizar o GoMock no seu projeto E_Bank.

#### 1. Instalando GoMock e Mockgen

Primeiro, adicione o GoMock e a ferramenta `mockgen` ao seu projeto:

```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@latest
```

A ferramenta `mockgen` é usada para gerar código de mock a partir das interfaces do seu projeto.

#### 2. Criando Mocks com GoMock

Vamos supor que você tenha uma interface `Store` no seu projeto que precisa ser mockada para testes:

**Interface Original (`store.go`):**

```go
package db

import "context"

type Store interface {
    GetAccount(ctx context.Context, id int64) (Account, error)
    DeleteAccount(ctx context.Context, id int64) error
}
```

Para gerar um mock dessa interface, execute o seguinte comando no terminal:

```bash
mockgen -source=db/store.go -destination=db/mock/store_mock.go -package=mockdb
```

Neste comando:

- **`-source=db/store.go`**: Caminho para o arquivo contendo a interface.
- **`-destination=db/mock/store_mock.go`**: Caminho onde o código de mock gerado será salvo.
- **`-package=mockdb`**: Nome do pacote onde o mock será gerado.

#### 3. Escrevendo Testes com GoMock

Depois de gerar o mock, você pode usá-lo em seus testes. Abaixo está um exemplo de como escrever um teste para a função que utiliza a interface `Store`:

**Exemplo de Teste (`store_test.go`):**

```go
package db_test

import (
    "context"
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/require"
    "github.com/usuario/e_bank/db/mock"
    "github.com/usuario/e_bank/db"
)

func TestDeleteAccount(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := mockdb.NewMockStore(ctrl)
    ctx := context.Background()
    accountID := int64(1)

    mockStore.EXPECT().
        DeleteAccount(ctx, accountID).
        Return(nil).
        Times(1)

    err := mockStore.DeleteAccount(ctx, accountID)
    require.NoError(t, err)
}
```

Neste exemplo:

- **`gomock.NewController(t)`**: Cria um novo controlador para gerenciar o ciclo de vida dos mocks.
- **`mockdb.NewMockStore(ctrl)`**: Cria uma nova instância do mock da interface `Store`.
- **`mockStore.EXPECT()`**: Define as expectativas para o mock, especificando que o método `DeleteAccount` deve ser chamado com os parâmetros fornecidos e deve retornar `nil`.
- **`Times(1)`**: Especifica que o método deve ser chamado exatamente uma vez.
- **`require.NoError(t, err)`**: Verifica se não houve erro ao chamar o método mockado.

#### 4. Executando os Testes

Após escrever os testes, você pode executá-los com o comando padrão do Go:

```bash
go test -v ./...
```

Isso executará todos os testes no seu projeto, incluindo aqueles que utilizam os mocks gerados pelo GoMock.

#### 5. Uso Avançado do GoMock

GoMock permite definir comportamentos complexos e encadeados, além de validações sobre os parâmetros passados para os métodos mockados. Aqui estão algumas funcionalidades adicionais:

- **Qualquer valor (`gomock.Any()`)**: Usado para corresponder a qualquer valor passado como argumento.
- **Sequências (`InOrder()`)**: Usado para garantir que as expectativas sejam satisfeitas em uma ordem específica.

**Exemplo de Uso Avançado:**

```go
mockStore.EXPECT().
    DeleteAccount(ctx, gomock.Any()).
    Return(nil).
    Times(1)
```

Este exemplo define que o método `DeleteAccount` deve ser chamado com qualquer valor de `id`.

---

Este tutorial fornece uma introdução ao uso do GoMock no seu projeto Go,
mostrando como ele pode facilitar a criação de testes de unidade robustos e isolados.
Com ele, você pode simular o comportamento de interfaces e verificar se o código funciona corretamente em diferentes cenários.

---

### Usando Bcrypt para Hashing de Senhas

`bcrypt` é uma biblioteca popular para hashing de senhas,
amplamente utilizada para garantir a segurança das credenciais dos usuários em aplicações web.
Ele aplica um algoritmo de hashing que é lento por design,
dificultando ataques de força bruta.
Neste tutorial, você aprenderá como integrar e utilizar
`bcrypt` no seu projeto E_Bank para proteger as senhas dos usuários.

#### 1. Instalando a Biblioteca Bcrypt

Para começar, você precisa instalar a biblioteca `bcrypt` no seu projeto Go:

```bash
go get golang.org/x/crypto/bcrypt
```

#### 2. Hashing de Senhas

Para armazenar as senhas dos usuários de forma segura,
você deve hashá-las antes de salvar no banco de dados.
Abaixo está um exemplo de como criar um hash de uma senha usando `bcrypt`:

```go
package util

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword recebe uma senha e retorna o hash bcrypt dessa senha.
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}
```

Neste exemplo:

- **`bcrypt.GenerateFromPassword`**: Gera um hash bcrypt a partir de uma senha.
  O segundo argumento (`bcrypt.DefaultCost`) define o custo de processamento do algoritmo,
  determinando quão intensivo em termos de CPU será o processo de hashing.
  O valor padrão (`bcrypt.DefaultCost`) é uma boa escolha para a maioria dos casos.

#### 3. Verificando Senhas

Quando um usuário tenta fazer login,
você precisa comparar a senha fornecida com o hash armazenado no banco de dados.
Abaixo está um exemplo de como realizar essa verificação:

```go
package util

import (
    "golang.org/x/crypto/bcrypt"
)

// CheckPassword compara uma senha com um hash bcrypt e retorna nil se coincidem.
func CheckPassword(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
```

Neste exemplo:

- **`bcrypt.CompareHashAndPassword`**: Compara o hash armazenado no banco de dados com a senha fornecida pelo usuário.
- Se as senhas coincidem, retorna `nil`; caso contrário, retorna um erro.

#### 4. Usando Bcrypt no Projeto

Agora que você tem as funções para hash e verificação,
pode usá-las em seu código de cadastro e autenticação de usuários.
Veja como pode ser feito:

**Exemplo de Cadastro de Usuário:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/usuario/e_bank/util"
)

func main() {
    password := "senhaSegura123"

    hashedPassword, err := util.HashPassword(password)
    if err != nil {
        log.Fatal("Erro ao criar hash da senha:", err)
    }

    fmt.Println("Senha Hasheada:", hashedPassword)
}
```

**Exemplo de Verificação de Senha:**

```go
package main

import (
    "fmt"
    "log"

    "github.com/usuario/e_bank/util"
)

func main() {
    password := "senhaSegura123"
    hashedPassword := "$2a$10$N9qo8uLOickgx2ZMRZo5e.SfIoB.z7x6WjVrqjRCJNNFPq5H7KJW6" // Exemplo de hash armazenado

    err := util.CheckPassword(password, hashedPassword)
    if err != nil {
        log.Fatal("Senha incorreta")
    }

    fmt.Println("Senha correta")
}
```

#### 5. Protegendo as Senhas com Bcrypt

Ao usar `bcrypt`, as senhas dos seus usuários estarão muito mais seguras contra ataques de força bruta.
No entanto, aqui estão algumas práticas recomendadas adicionais:

- **Sempre hash as senhas antes de armazená-las**: Nunca armazene senhas em texto plano.
- **Ajuste o custo conforme necessário**: O custo (`bcrypt.DefaultCost`) pode ser ajustado para aumentar a segurança,
  mas tenha em mente que valores muito altos podem impactar a performance da aplicação.
- **Verifique as senhas com segurança**: Sempre use `bcrypt.CompareHashAndPassword` para comparar senhas,
  nunca compare hashes diretamente.

---

Este tutorial explica o que é o `bcrypt` e como você pode usá-lo no seu projeto para proteger as senhas dos usuários de forma eficaz e segura.

---

## Autenticação Baseada em Token com PASETO

[PASETO (Platform-Agnostic Security Tokens)](https://paseto.io/) é uma alternativa mais segura e moderna ao JWT (JSON Web Tokens) para autenticação baseada em tokens. Ao contrário do JWT, que tem várias implementações inseguras, PASETO foi projetado para ser seguro por padrão, evitando potenciais vulnerabilidades conhecidas.

Neste tutorial, vamos aprender o que é PASETO e como implementá-lo no projeto E_Bank para autenticação de usuários.

**1. O que é PASETO?**
PASETO é um formato de token assinado e/ou criptografado que, ao ser gerado, garante:

- **Autenticidade:** O token não pode ser alterado sem ser detectado.
- **Segurança:** PASETO previne vulnerabilidades como a confusão de algoritmos, que pode ocorrer com JWT.
- **Simplicidade:** Comparado ao JWT, PASETO é mais simples de usar corretamente, reduzindo o risco de erros na implementação.
  Existem dois tipos de PASETO:

- V1 (Assinatura com RSA) e V2 (Assinatura com Ed25519 e criptografia com XChaCha20-Poly1305). Neste exemplo, usaremos a versão V2, que é recomendada para aplicações modernas.

**2. Instalando a Biblioteca PASETO**
Primeiro, instale a biblioteca Go para trabalhar com PASETO:

```bash
go get github.com/o1egl/paseto
```

**3. Implementando a Geração e Verificação de Tokens PASETO**
Aqui estão os passos para implementar PASETO no projeto E_Bank:

Passo 1: Criar uma chave secreta

PASETO usa uma chave secreta para assinar os tokens. A chave secreta deve ter pelo menos 32 bytes de comprimento.

```go
package util

import (
    "time"
    "github.com/o1egl/paseto"
    "golang.org/x/crypto/chacha20poly1305"
    "fmt"
)

type PasetoMaker struct {
    paseto       *paseto.V2
    symmetricKey []byte
}

// NewPasetoMaker cria uma nova instância de PasetoMaker com uma chave secreta de pelo menos 32 bytes.
func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
    if len(symmetricKey) != chacha20poly1305.KeySize {
        return nil, fmt.Errorf("tamanho inválido da chave secreta: deve ter %d bytes", chacha20poly1305.KeySize)
    }

    maker := &PasetoMaker{
        paseto:       paseto.NewV2(),
        symmetricKey: []byte(symmetricKey),
    }
    return maker, nil
}

```

Passo 2: Geração de Tokens

Agora, vamos criar a função para gerar tokens PASETO com um tempo de expiração e dados customizados (payload), como o ID do usuário.

```go
type TokenPayload struct {
    Username string    `json:"username"`
    IssuedAt time.Time `json:"issued_at"`
    ExpiredAt time.Time `json:"expired_at"`
}

// CreateToken gera um novo token PASETO com um tempo de expiração.
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *TokenPayload, error) {
    payload := &TokenPayload{
        Username:  username,
        IssuedAt:  time.Now(),
        ExpiredAt: time.Now().Add(duration),
    }

    token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
    return token, payload, err
}

```

Passo 3: Verificação de Tokens

Depois de gerar o token, precisamos de uma função para verificar sua autenticidade e validade:

```go
// VerifyToken verifica se o token PASETO é válido e retorna o payload.
func (maker *PasetoMaker) VerifyToken(token string) (*TokenPayload, error) {
    var payload TokenPayload
    err := maker.paseto.Decrypt(token, maker.symmetricKey, &payload, nil)
    if err != nil {
        return nil, fmt.Errorf("token inválido: %v", err)
    }

    if time.Now().After(payload.ExpiredAt) {
        return nil, fmt.Errorf("token expirado")
    }

    return &payload, nil
}

```

**4. Integrando PASETO no Fluxo de Autenticação**
Aqui está como integrar o PASETO ao fluxo de autenticação do seu projeto:

**1. Criação do Token na Autenticação:** Quando o usuário faz login com credenciais corretas, você gera um token PASETO com uma duração específica.

```go
package main

import (
    "fmt"
    "time"
    "log"
    "github.com/usuario/e_bank/util"
)

func main() {
    // Crie uma chave secreta (32 bytes)
    secretKey := "minha-chave-super-segura-32bytes!"

    pasetoMaker, err := util.NewPasetoMaker(secretKey)
    if err != nil {
        log.Fatal("Erro ao criar Paseto Maker:", err)
    }

    token, payload, err := pasetoMaker.CreateToken("usuario123", time.Hour)
    if err != nil {
        log.Fatal("Erro ao criar token:", err)
    }

    fmt.Println("Token:", token)
    fmt.Printf("Payload: %+v\n", payload)
}

```

**2. Verificação do Token em Endpoints Protegidos:**
Em endpoints que exigem autenticação,
como operações bancárias ou consulta de dados,
você verifica o token enviado pelo cliente.

```go
package main

import (
    "fmt"
    "log"
    "github.com/usuario/e_bank/util"
)

func main() {
    secretKey := "minha-chave-super-segura-32bytes!"

    pasetoMaker, err := util.NewPasetoMaker(secretKey)
    if err != nil {
        log.Fatal("Erro ao criar Paseto Maker:", err)
    }

    token := "token_gerado_aqui"

    payload, err := pasetoMaker.VerifyToken(token)
    if err != nil {
        log.Fatal("Erro ao verificar token:", err)
    }

    fmt.Printf("Payload válido: %+v\n", payload)
}

```

**5. Práticas Recomendadas**

**- Use uma chave secreta segura:** Certifique-se de que a chave secreta tenha pelo menos 32 bytes e que seja gerada de forma segura.
**- Defina um tempo de expiração apropriado:** Tokens devem expirar após um período adequado para evitar problemas de segurança em caso de vazamento.
**- Use HTTPS:** Sempre transmita tokens em conexões seguras para evitar que sejam interceptados.

---

Com PASETO, você pode implementar autenticação baseada em tokens de forma segura no E_Bank. Este tutorial cobre a criação e verificação de tokens, além de fornecer as bases para integrar PASETO ao fluxo de autenticação do seu projeto.
