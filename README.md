# Adquirer

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)

API de adquirencia para tokenizacao e autorizacao de transacoes com cartao de credito, desenvolvida em Go.

## Funcionalidades

- Tokenizacao de cartao com geracao de token HMAC-SHA256 em Base64.
- Identificacao da bandeira do cartao pelo numero informado.
- Retorno padronizado com metadados do cartao (brand, cardToken, expirationMonth, expirationYear e lastFourDigits).
- Fluxo de autorizacao por token.
- Validacao de cartao com algoritmo de Luhn.

## Tecnologias

- Go 1.24.2
- MySQL (go-sql-driver/mysql)
- godotenv

## Estrutura (resumo)

```
adquirer/
├── db/
├── entity/
│   ├── authorization-request.go
│   ├── card.go
│   ├── validate-request.go
│   └── validate-response.go
├── handler/
│   └── handler.go
├── model/
│   ├── get-card-flag.go
│   ├── get-card.go
│   ├── get-token.go
│   └── update-card.go
├── repository/
├── validation/
├── main.go
├── go.mod
└── README.md
```

## Configuracao

Variaveis de ambiente esperadas:

```env
DB_NAME=adquirer
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
CARD_TOKEN_SECRET=sua_chave_secreta
```

## Como executar

```bash
go run .
```

Servidor padrao: porta 8081.

## Endpoints

### POST /adquirer/valid-token

Tokeniza o cartao e retorna metadados.

Request body:

```json
{
  "card": "4111111111112454",
  "cvv": "123",
  "venc": "12/2026",
  "holderName": "Nome Cliente",
  "expirationMonth": "12",
  "expirationYear": "2026",
  "total": 100.50
}
```

Response de sucesso:

```json
{
  "message": "Tokenrização realizadao com sucesso",
  "code": "00",
  "brand": "Visa",
  "cardToken": "QqGzPLHZHMg/dE5RYQBh4aomTueENoE4hrZyLg8YSeQ=",
  "expirationMonth": "12",
  "expirationYear": "2026",
  "lastFourDigits": "2454"
}
```

Response de erro (exemplo):

```json
{
  "message": "Erro ao gerar token",
  "code": "96",
  "brand": "",
  "cardToken": "",
  "expirationMonth": "12",
  "expirationYear": "2026",
  "lastFourDigits": "2454"
}
```

### POST /adquirer/authorization

Autoriza a transacao a partir de um token previamente gerado.

Request body:

```json
{
  "token": "QqGzPLHZHMg/dE5RYQBh4aomTueENoE4hrZyLg8YSeQ=",
  "amount": 50.0
}
```

Resposta de sucesso:

```json
{
  "message": "Autorização realizada com sucesso",
  "code": "00",
  "brand": "",
  "cardToken": "",
  "expirationMonth": "",
  "expirationYear": "",
  "lastFourDigits": ""
}
```

### POST /adquirer/capture

Endpoint de captura mantido no projeto para fluxo de atualizacao de saldo.

## Regras de validacao

- O numero do cartao e validado com algoritmo de Luhn.
- A bandeira e inferida no backend com base em prefixo e tamanho do cartao.
- O token e gerado com HMAC-SHA256 + Base64 usando CARD_TOKEN_SECRET.
