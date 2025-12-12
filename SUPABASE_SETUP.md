# Guia de Configuração - Supabase PostgreSQL

## Passo 1: Obter Credenciais do Supabase

1. Acesse: https://app.supabase.com
2. Selecione seu projeto
3. Vá em **Settings** > **Database**
4. Copie os dados de conexão:
   - **Host**: db.dcdakvwglegcqkozgawy.supabase.co
   - **Port**: 5432
   - **Database**: postgres
   - **User**: postgres
   - **Password**: [Sua senha]

## Passo 2: Configurar .env

Edite o arquivo `.env` na raiz do projeto:

```env
# Database Configuration - Supabase
DB_HOST=db.dcdakvwglegcqkozgawy.supabase.co
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=INSIRA_SUA_SENHA_AQUI
DB_NAME=postgres
DB_SSLMODE=require
DB_TIMEZONE=UTC

# Server Configuration
SERVER_PORT=8080
SERVER_ENV=production

# JWT Configuration
JWT_SECRET=ALTERE_PARA_UMA_CHAVE_SEGURA
JWT_EXPIRATION=24h

# CORS Configuration
CORS_ORIGIN=https://pantheonconcursos.com.br,http://localhost:3000
```

## Passo 3: Testar Conexão

Execute o comando de teste:

```bash
go run cmd/test-db/main.go
```

Você deve ver:
```
✅ Successfully connected to Supabase PostgreSQL!
✅ Database migration completed!
✅ Test query result: 2025-12-12 ...
```

## Passo 4: Iniciar a API

```bash
go run cmd/api/main.go
```

## Troubleshooting

### Erro: "connectex: Uma tentativa de conexão falhou"

**Solução**: 
- Verifique se a senha está correta
- Verifique a conectividade de rede (firewall/VPN)
- Certifique-se de que o projeto Supabase está ativo
- Tente a conexão via psql:
  ```bash
  psql -h db.dcdakvwglegcqkozgawy.supabase.co -U postgres -d postgres
  ```

### Erro: "password authentication failed"

**Solução**:
- Resete a senha no Supabase Dashboard
- Copie a nova senha no arquivo `.env`
- Tente novamente

## Variáveis de Ambiente Importantes

| Variável | Descrição |
|----------|-----------|
| DB_HOST | Host do Supabase PostgreSQL |
| DB_PORT | Porta (sempre 5432) |
| DB_USER | Usuário (geralmente postgres) |
| DB_PASSWORD | Senha (obtenha no Dashboard) |
| DB_NAME | Banco de dados (postgres ou personalizado) |
| DB_SSLMODE | Sempre `require` para Supabase |
| SERVER_ENV | `development` ou `production` |
| JWT_SECRET | Chave secreta para tokens (mude em produção) |

## Segurança

⚠️ **IMPORTANTE**:
- Nunca commit o arquivo `.env` no git (já está no `.gitignore`)
- Use variáveis de ambiente em produção
- Altere a senha do banco regularmente
- Altere o JWT_SECRET para um valor aleatório e seguro

## Comandos Úteis

```bash
# Testar conexão
go run cmd/test-db/main.go

# Executar a API
go run cmd/api/main.go

# Build para produção
go build -o bin/api cmd/api/main.go

# Executar testes
go test ./...
```
