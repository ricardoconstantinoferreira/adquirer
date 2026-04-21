# Adquirer

API de adquirência para validação e autorização de transações com cartão de crédito, desenvolvida em Go.

## Tecnologias

- **Go** 1.24.2
- **MySQL** — via `go-sql-driver/mysql` v1.9.3
- **godotenv** v1.5.1 — carregamento de variáveis de ambiente

## Estrutura do projeto

```
adquirer/
├── db/
│   └── db.go            # Conexão com o banco de dados MySQL
├── handler/
│   └── handler.go       # Handler HTTP do endpoint de validação
├── validation/
│   └── validation.go    # Algoritmo de validação de Luhn
├── main.go              # Bootstrap do servidor HTTP (porta 8081)
├── .env                 # Credenciais do banco (não versionado)
├── .env.example         # Template de variáveis de ambiente
└── go.mod
```

## Configuração

Copie o arquivo `.env.example` para `.env` e preencha com suas credenciais:

```bash
cp .env.example .env
```

Variáveis esperadas no `.env`:

```env
DB_NAME=adquirer
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
```

## Como executar

```bash
go run .
```

O servidor inicia na porta **8081**.

## Endpoints

### `POST /adquirer/valid`

Valida um cartão de crédito aplicando o algoritmo de Luhn e retorna o código de resposta da transação.

**Request Body**

```json
{
    "card": "4242 4242 4242 4242",
    "cvv": "435",
    "venc": "10/90",
    "total": 54.99
}
```

| Campo   | Tipo    | Descrição                             |
|---------|---------|---------------------------------------|
| `card`  | string  | Número do cartão (com ou sem espaços) |
| `cvv`   | string  | Código de segurança                   |
| `venc`  | string  | Data de vencimento (MM/AA)            |
| `total` | float64 | Valor da transação                    |

**Respostas**

| Código | Mensagem                         | Descrição                         |
|--------|----------------------------------|-----------------------------------|
| `00`   | Transacao autorizada com sucesso | Cartão válido, transação aprovada |
| `14`   | Cartão inválido                  | Número do cartão inválido (Luhn)  |
| `96`   | Payload inválido                 | JSON malformado ou ausente        |

**Exemplo — cartão válido**

```bash
curl -X POST http://localhost:8081/adquirer/valid \
  -H "Content-Type: application/json" \
  -d '{"card": "4242 4242 4242 4242", "cvv": "435", "venc": "10/90", "total": 54.99}'
```

```json
{
    "message": "Transacao autorizada com sucesso",
    "code": "00"
}
```

**Exemplo — cartão inválido**

```bash
curl -X POST http://localhost:8081/adquirer/valid \
  -H "Content-Type: application/json" \
  -d '{"card": "1234 5678 9012 3456", "cvv": "123", "venc": "01/25", "total": 10.00}'
```

```json
{
    "message": "Cartão inválido",
    "code": "14"
}
```

## Validação de Luhn

O número do cartão é validado pelo [algoritmo de Luhn](https://en.wikipedia.org/wiki/Luhn_algorithm), que detecta erros de digitação e números inválidos. Espaços no número do cartão são ignorados automaticamente.
