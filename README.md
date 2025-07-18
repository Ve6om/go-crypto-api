# go-crypto-api

A very simple Go REST API that fetches live crypto prices. Supports choosing different currencies and time zones.


## How to run

Run the server:

```bash
go run main.go
```

## How to request

```
GET /price?coin=bitcoin&currency=usd&timezone=America/New_York
```
### Response
```json
{
  "coin": "bitcoin",
  "currency": "usd",
  "price": 29874.53,
  "timestamp": "2025-07-18T12:34:56-04:00"
}
```