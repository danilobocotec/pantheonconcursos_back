# Como Acessar o Swagger da API

## Acesso Rápido

Após iniciar o servidor, acesse:

**URL do Swagger UI:**
```
http://localhost:8080/swagger/index.html
```

## Como Iniciar o Servidor

### Ambiente de Desenvolvimento (Banco Local)
```powershell
$env:SERVER_ENV="development"
go run cmd/api/main.go
```

### Ambiente de Produção (Supabase)
```powershell
$env:SERVER_ENV="production"
go run cmd/api/main.go
```

## Endpoints Documentados

O Swagger documenta todos os endpoints da API:

### 🔐 Autenticação (`/api/v1/auth`)
- **POST /auth/register** - Registrar novo usuário
  - Campos: `full_name`, `email`, `password`, `confirm`
- **POST /auth/login** - Login de usuário
  - Campos: `email`, `password`
- **POST /auth/refresh** - Atualizar token JWT (requer autenticação)

### 👤 Usuários (`/api/v1/users`)
- **GET /users** - Listar usuários (paginação)
- **GET /users/:id** - Obter usuário por ID
- **POST /users** - Criar usuário
- **PUT /users/:id** - Atualizar usuário
- **DELETE /users/:id** - Deletar usuário

### 💚 Health (`/api/v1/health`)
- **GET /health** - Verificar status da API

## Como Usar Autenticação no Swagger

1. **Registre um usuário** via `/auth/register`
2. **Faça login** via `/auth/login` e copie o `token` da resposta
3. **Clique no botão "Authorize"** (cadeado) no topo da página
4. **Digite**: `Bearer SEU_TOKEN_AQUI`
5. **Clique em "Authorize"**
6. Agora pode testar endpoints protegidos

## Exemplo de Registro

```json
{
  "full_name": "João Silva",
  "email": "joao@example.com",
  "password": "SenhaForte123",
  "confirm": "SenhaForte123"
}
```

## Exemplo de Login

```json
{
  "email": "joao@example.com",
  "password": "SenhaForte123"
}
```

## Regenerar Documentação

Se adicionar novos endpoints ou alterar anotações:

```powershell
swag init -g cmd/api/main.go
```

## Arquivos Gerados

- `docs/docs.go` - Código Go da documentação
- `docs/swagger.json` - Especificação OpenAPI em JSON
- `docs/swagger.yaml` - Especificação OpenAPI em YAML

## URLs Alternativas

- **JSON da especificação**: http://localhost:8080/swagger/doc.json
- **YAML da especificação**: http://localhost:8080/swagger/swagger.yaml

## Troubleshooting

### Swagger não carrega
- Certifique-se de que o servidor está rodando
- Verifique se acessou a URL correta com `/index.html`
- Confirme que os arquivos em `docs/` existem

### Endpoints não aparecem
- Execute `swag init -g cmd/api/main.go`
- Reinicie o servidor
- Verifique se as anotações estão corretas nos handlers

### Autenticação não funciona
- Use o formato: `Bearer SEU_TOKEN`
- Certifique-se de incluir o espaço após "Bearer"
- Verifique se o token não expirou (validade: 24h)
