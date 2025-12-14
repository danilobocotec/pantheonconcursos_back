# ThePantheon API

Uma API REST robusta desenvolvida em Go Lang com arquitetura em camadas.

## Características

- ✅ Arquitetura em camadas (Handler, Service, Repository)
- ✅ Autenticação JWT
- ✅ CRUD de Usuários
- ✅ Middlewares customizados (CORS, Logging, Autenticação)
- ✅ Banco de dados PostgreSQL com GORM
- ✅ Validação de entrada
- ✅ Tratamento de erros

## Pré-requisitos

- Go 1.21+
- PostgreSQL 12+
- Git

## Instalação

1. Clone o repositório ou navegue até o diretório do projeto

2. Instale as dependências:
```bash
go mod download
go mod tidy
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

4. Atualize o `.env` com suas configurações:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=thepantheon_db
SERVER_PORT=8080
JWT_SECRET=sua_chave_secreta
```

5. Crie o banco de dados PostgreSQL:
```bash
createdb thepantheon_db
```

## Execução

Inicie o servidor:
```bash
go run cmd/api/main.go
```

O servidor será iniciado em `http://localhost:8080`

## Estrutura do Projeto

```
.
├── cmd/
│   └── api/
│       └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── config/               # Configurações da aplicação
│   ├── handler/              # Handlers HTTP
│   ├── service/              # Lógica de negócio
│   ├── repository/           # Acesso a dados
│   └── model/                # Modelos de dados
├── pkg/
│   ├── middleware/           # Middlewares da API
│   └── errors/               # Tratamento de erros
├── migrations/               # Migrações do banco de dados
├── go.mod                    # Dependências do Go
├── go.sum                    # Checksums das dependências
└── README.md                 # Esta documentação
```

## Endpoints da API

### Health Check
- `GET /api/v1/health`

### Usuários
- `POST /api/v1/users` - Criar usuário
- `GET /api/v1/users` - Listar usuários
- `GET /api/v1/users/:id` - Obter usuário por ID
- `PUT /api/v1/users/:id` - Atualizar usuário
- `DELETE /api/v1/users/:id` - Deletar usuário

### Autenticação
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/register` - Registrar novo usuário
- `POST /api/v1/auth/refresh` - Atualizar token JWT

## Exemplos de Requisições

### Registrar Usuário
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe",
    "password": "securepassword123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword123"
  }'
```

### Obter Usuário
```bash
curl -X GET http://localhost:8080/api/v1/users/UUID \
  -H "Authorization: Bearer SEU_TOKEN_JWT"
```

## Dependências Principais

- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL Driver** - Driver PostgreSQL para Go
- **JWT** - Autenticação com tokens JWT
- **UUID** - Geração de UUIDs
- **Godotenv** - Carregamento de variáveis de ambiente

## Desenvolvimento

### Adicionar Novo Modelo

1. Crie o modelo em `internal/model/`
2. Crie o repositório em `internal/repository/`
3. Crie o serviço em `internal/service/`
4. Crie o handler em `internal/handler/`
5. Registre as rotas em `cmd/api/main.go`

### Executar Testes

```bash
go test ./...
```

### Build para Produção

```bash
go build -o bin/api cmd/api/main.go
```

## Contribuindo

1. Crie um branch para sua feature (`git checkout -b feature/AmazingFeature`)
2. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
3. Push para o branch (`git push origin feature/AmazingFeature`)
4. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## Suporte

Para suporte, abra uma issue no repositório ou entre em contato com o time de desenvolvimento.
