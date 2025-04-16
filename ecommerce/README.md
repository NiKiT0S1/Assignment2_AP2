# E-Commerce Platform (Microservices)

## Overview

This project is a basic e-commerce platform built with Go, structured using Clean Architecture and composed of three microservices:

- API Gateway – routes incoming requests to the appropriate service, handles logging and basic authentication.

- Inventory Service – manages product data, categories, stock, and prices.

- Order Service – handles order creation, status updates, and product quantities per order.

All services are written in Go using the Gin framework and connected to PostgreSQL for persistence.

## Project Structure

```
.
├── inventory/              # Inventory microservice
│   ├── cmd/
│   └── internal/
├── order/                 # Order microservice
│   ├── cmd/
│   └── internal/
└── gateway/               # API Gateway
├── cmd/
└── internal/
```

## Requirements

- Go 1.18 or higher

- PostgreSQL

- Basic understanding of RESTful APIs

## Installation

### 1. Install Go Dependencies:

   In each service directory:
   ```
    cd inventoryService
    go mod tidy

    cd ../orderService
    go mod tidy

    cd ../apiGateway
    go mod tidy
   ```

### 2. Set Up PostgreSQL Tables:
Inventory Service:
```
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    stock INT NOT NULL,
    category_id INT NOT NULL
);
```

Order Service:
```
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL,
    quantity INT NOT NULL
);
```

### 3. Running the services:
Inventory Service:
```
cd inventoryService
go run cmd/main.go
```
Runs on http://localhost:8081

Order Service:
```
cd orderService
go run cmd/main.go
```
Runs on http://localhost:8082

API-Gateway:
```
cd apiGateway
go run cmd/main.go
```
Runs on http://localhost:8080

## Authentication
All API Gateway endpoints are protected by Basic Auth:

- Username: admin

- Password: 1234

Base64 encoded: YWRtaW46MTIzNA==

Include this in your Authorization header:
```
Authorization: Basic YWRtaW46MTIzNA==
```

## Sample API Requests (via API Gateway)

### Create Product:
```
curl -X POST http://localhost:8080/products \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{"name": "iPhone", "description": "Apple smartphone", "price": 999.99, "stock": 10, "category_id": 1}'
```

### List Products:
```
curl -X GET http://localhost:8080/products \
  -H "Authorization: Basic YWRtaW46MTIzNA=="
```

### Create Order:
```
curl -X POST http://localhost:8080/orders \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {"product_id": 1, "quantity": 2}
    ]
  }'
```

### Get Order by ID:
```
curl -X GET http://localhost:8080/orders/1 \
  -H "Authorization: Basic YWRtaW46MTIzNA=="
```

### Update Order Status:
```
curl -X PATCH http://localhost:8080/orders/1 \
  -H "Authorization: Basic YWRtaW46MTIzNA==" \
  -H "Content-Type: application/json" \
  -d '{"status": "completed"}'
```

## Logging
Each request is logged to the console with:

- IP address

- Method

- Endpoint

- Status code

- Duration


## HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Create successful
- `400 Bad Request`: Missing required parameters or invalid input
- `401 Unauthorized (Auth)`: Unauthorized

- `404 Not Found`: No news found for the specified cryptocurrency
- `500 Internal Server Error`: Server-side error

## Notes
- All API calls must go through the API Gateway (localhost:8080)

- Services run independently; no Docker or gRPC is used

- You can easily extend this to include user authentication, payment providers, etc.