
# Portal-BI Backend

Este é o backend do projeto **Portal-BI**, desenvolvido em Go. Ele fornece uma API RESTful para gerenciar os dados do sistema, utilizando as bibliotecas **Gin**, **SQLC** e **Swag** para facilitar o desenvolvimento e a documentação.

---

## 🛠️ Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para o desenvolvimento.
- **Gin**: Framework minimalista para criar APIs RESTful rápidas e eficientes.
- **SQLC**: Gera código Go diretamente das suas queries SQL, simplificando a interação com o banco de dados.
- **Swag**: Gera automaticamente a documentação da API no formato Swagger.

---

## 🔧 Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/) (opcional, para usar com contêineres)
- [PostgreSQL](https://www.postgresql.org/) ou outro banco de dados compatível.

---

## 🚀 Como Executar o Projeto

1. **Clone o Repositório**:

   ```bash
   git clone https://github.com/seu-usuario/portal-bi-backend.git
   cd portal-bi-backend
   ```

2. **Instale as Dependências**:

   ```bash
   go mod tidy
   ```

3. **Configure o Banco de Dados**:

   Crie um banco de dados e configure as variáveis de ambiente no arquivo `.env`:

   ```env
   DB_DRIVER=postgres
   DB_SOURCE=postgresql://usuario:senha@localhost:5432/portal_bi?sslmode=disable
   ```

4. **Gere o Código do SQLC**:

   Execute o seguinte comando para gerar os arquivos necessários para o acesso ao banco de dados:

   ```bash
   sqlc generate
   ```

5. **Execute o Projeto**:

   ```bash
   go run main.go
   ```

6. **Acesse a API**:

   Por padrão, o servidor estará disponível em: `http://localhost:8080`.

---

## 📖 Rotas Disponíveis

### 📄 GET
- **`/api/v1/resource`**: Retorna a lista de recursos.
- **`/api/v1/resource/:id`**: Retorna um recurso específico pelo ID.

### ✏️ POST
- **`/api/v1/resource`**: Cria um novo recurso.

### 🛠️ PUT
- **`/api/v1/resource/:id`**: Atualiza os dados de um recurso existente.

### ❌ DELETE
- **`/api/v1/resource/:id`**: Remove um recurso pelo ID.

---

## 🧩 Estrutura do Projeto

```
portal-bi-backend/
│
├── api/               # Rotas e controladores
├── db/                # Configurações do banco de dados (SQLC)
│   ├── queries/       # Queries SQL gerenciadas pelo SQLC
│   └── migrations/    # Arquivos de migração do banco
├── docs/              # Documentação Swagger gerada pelo Swag
├── config/            # Configurações gerais (ex.: variáveis de ambiente)
├── models/            # Estrutura de dados e modelos
├── utils/             # Funções auxiliares
└── main.go            # Ponto de entrada da aplicação
```

---

## 📄 Documentação Swagger

A documentação da API é gerada automaticamente utilizando o **Swag**. Após executar o servidor, acesse:

**`http://localhost:8080/swagger/index.html`**

Para atualizar a documentação:

1. Certifique-se de que o Swag está instalado:

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. Gere os arquivos de documentação:

   ```bash
   swag init
   ```

---

## 🐋 Executando com Docker

Um arquivo `Dockerfile` está incluído no projeto. Para executar via Docker:

1. **Construa a imagem**:

   ```bash
   docker build -t portal-bi-backend .
   ```

2. **Execute o contêiner**:

   ```bash
   docker run -p 8080:8080 --env-file .env portal-bi-backend
   ```

---

## 📝 Contribuições

Sinta-se à vontade para contribuir com melhorias, correções de bugs ou novas funcionalidades. Abra uma [issue](https://github.com/seu-usuario/portal-bi-backend/issues) ou envie um [pull request](https://github.com/seu-usuario/portal-bi-backend/pulls).

---

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.

---

## ✨ Autor

Desenvolvido com 💻 por [Seu Nome](https://github.com/seu-usuario).
