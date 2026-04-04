# Starti Backend

API REST desenvolvida em Go com Gin, GORM e PostgreSQL.

## Stack

- **Go 1.22+**
- **Gin** — framework HTTP
- **GORM** — ORM
- **PostgreSQL** — banco de dados
- **Docker / Docker Compose** — conteinerização

## Arquitetura

```
cmd/api/          → entrypoint
internal/
  config/         → variáveis de ambiente
  db/             → conexão e migrations
  model/          → entidades (User, Post, Comment)
  repositories/   → acesso ao banco (interface + implementação)
  services/       → regras de negócio
  handlers/       → controllers HTTP
  router/         → definição de rotas
```

## Rodando localmente

### 1. Clonar e configurar

```bash
git clone https://github.com/JoaoGSantiago/starti-backend.git
cd starti-backend
cp .env.example .env
```

### 2. Subir apenas o banco com Docker

```bash
docker compose up db -d
```

### 3. Rodar a API

```bash
go run ./cmd/api
```

### Ou rodar tudo com Docker

```bash
docker compose up --build
```

A API ficará disponível em `http://localhost:8080`.

---

## Endpoints

Base URL: `/api/v1`

### Users

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/users` | Listar todos os usuários |
| POST | `/users` | Criar usuário |
| GET | `/users/:id` | Buscar usuário por ID |
| PUT | `/users/:id` | Atualizar usuário |
| DELETE | `/users/:id` | Apagar usuário |
| GET | `/users/:id/posts` | Listar posts públicos do usuário |
| GET | `/users/:id/comments` | Listar comentários do usuário em posts públicos |

### Posts

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/posts` | Criar publicação |
| GET | `/posts/:id` | Buscar publicação por ID |
| PUT | `/posts/:id` | Atualizar publicação |
| DELETE | `/posts/:id` | Apagar publicação |
| PATCH | `/posts/:id/archive` | Arquivar publicação |
| GET | `/posts/:id/comments` | Listar comentários da publicação |

### Comments

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/comments` | Criar comentário |
| PUT | `/comments/:id` | Atualizar comentário |
| DELETE | `/comments/:id` | Apagar comentário |

---

## Exemplos de uso

### 1. Criar usuário (rota pública)

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "joao",
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "secret123",
    "biography": "Desenvolvedor Go"
  }'
```

### 2. Login — obter o token

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "joao@example.com", "password": "secret123"}'
```

Resposta:
```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." }
```

Salve o token numa variável para facilitar os próximos comandos:

```bash
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### 3. Criar post

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "text": "Meu primeiro post!"}'
```

### 4. Criar comentário

```bash
curl -X POST http://localhost:8080/api/v1/comments \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "post_id": 1, "message": "Ótimo post!"}'
```

### 5. Arquivar post

```bash
curl -X PATCH http://localhost:8080/api/v1/posts/1/archive \
  -H "Authorization: Bearer $TOKEN"
```

### 6. Listar posts públicos de um usuário

```bash
curl http://localhost:8080/api/v1/users/1/posts \
  -H "Authorization: Bearer $TOKEN"
```
