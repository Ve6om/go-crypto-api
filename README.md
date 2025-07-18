# go-crypto-api

A very simple Go REST API that fetches live crypto prices. Supports choosing different currencies and time zones.


## How to Run

1. Clone the repository and enter the directory.

2. (Optional) Create a `.env` file in the root of the project to customize settings:

```env
# The frontend URL to allow CORS requests from (e.g. your React app)
# If not set, CORS will be disabled.
FRONTEND_URL=http://localhost:5173

# The port the API should listen on
# If not set, the default is 8080
PORT=8080
```

3. Run the server:
```bash
go run main.go
```

# Endpoints
> #### `GET /price`
Fetch the current price of a specific cryptocurrency.

### Query Parameters:
- `coin` (required): The coin ID (e.g., `bitcoin`). A full list of supported coin IDs can be found using the `/coins` endpoint shown below.
- `currency` (optional): The target currency (default: `usd`)
- `timezone` (optional): The time zone for the timestamp (default: UTC)

### Example Request:
```
GET /price?coin=bitcoin&currency=usd&timezone=America/New_York
```

### Example Response
```json
{
  "coin": "bitcoin",
  "currency": "usd",
  "price": 29874.53,
  "timestamp": "2025-07-18T12:34:56-04:00"
}
```



> #### `GET /coins`
Returns a list of the coins.

### Example Request:
`GET /coins`

```json
[
  {
    "id": "bitcoin",
    "symbol": "btc",
    "name": "Bitcoin"
  },
  {
    "id": "ethereum",
    "symbol": "eth",
    "name": "Ethereum"
  },
  {
    "id": "ripple",
    "symbol": "xrp",
    "name": "XRP"
  }
]
```
