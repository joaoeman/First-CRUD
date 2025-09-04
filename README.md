1. First-CRUD (Go REST API)

First-CRUD é uma API REST simples desenvolvida em Golang, implementando operações CRUD (Create, Read, Update, Delete) básicas. Este projeto é ideal para quem está começando com Golang, servindo como template ou sandbox para aprender e expandir.

Tecnologias Utilizadas

Golang — linguagem principal, 100% do código e lógica da API. 
Postgres - Banco de dados utilizado para controlar dados publicados,removidos, lidos ou alterados.

 Estrutura típica de projetos Go, com pacotes organizados por responsabilidade:

  * `models/` — definição das entidades e seus atributos
  * `handlers/` — lógica dos endpoints da API
  * `main.go` — ponto de entrada, configuração das rotas e inicialização do servidor
  * `go.mod`, `go.sum` — gerenciamento de dependências


2. Estrutura do Projeto


First-CRUD/
 ├── models/       # Definições de modelos de dados
 ├── handlers/     # Endpoints e lógica de CRUD
 ├── main.go       # Inicialização da aplicação e roteamento
 ├── go.mod        # Módulo e dependências
 └── go.sum        # Versões exatas das dependências


3. Funcionalidades

* Create — Cria um novo item no sistema.
* Read — Recupera um registro ou lista de registros existentes.
* Update — Modifica dados de um item existente.
* Delete — Remove um item do sistema.

4. Como Executar:
    
    1. Clone o repositório:
    
       ```bash
       git clone https://github.com/joaoeman/First-CRUD.git
       cd First-CRUD
       ```
    
    2. Instale as dependências (se necessário):
    
       ```bash
       go mod tidy
       ```
    
    3. Execute a aplicação:
    
       ```bash
       go run main.go
       ```
    
    4. A API será iniciada localmente (por padrão, em `http://localhost:8080`) — você pode testar os endpoints via Postman, curl, ou outra ferramenta HTTP.
    

