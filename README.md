# E-Commerce Platform (Microservices with gRPC & Clean Architecture)

## Overview

This project is a simple e-commerce platform built using Go, structured according to Clean Architecture, and composed of several microservices communicating via gRPC. The platform includes:

- Inventory Service – manages product information, categories, stock, and pricing.

- Order Service – handles order creation, order status updates, and items per order.

- User Service – handles user registration, basic authentication, and profile retrieval.

- API Gateway (**currently under development**) – will serve as a central entry point for clients, providing routing, logging, and authentication.

Each service uses PostgreSQL for persistence and gRPC for inter-service communication.

## Project Structure

```
├── inventoryService/        # Inventory microservice (gRPC + PostgreSQL)
│   ├── cmd/
│   └── internal/
├── orderService/            # Order microservice (gRPC + PostgreSQL)
│   ├── cmd/
│   └── internal/
├── userService/             # User microservice (gRPC + PostgreSQL)
│   ├── cmd/
│   └── internal/
└── apiGateway/              # API Gateway (REST + gRPC, under construction)
    ├── cmd/
    └── internal/
```

## Requirements

- Go 1.18 or higher

- PostgreSQL

- Protocol Buffers + gRPC plugins:

  - protoc

  - protoc-gen-go

  - protoc-gen-go-grpc

- Postman (with gRPC support) or grpcurl

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

User Service
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);
```

### 3. Running the services:
```
# Run Inventory Service
cd inventoryService
go run cmd/main.go

# Run Order Service
cd orderService
go run cmd/main.go

# Run User Service
cd userService
go run cmd/main.go

# API Gateway (not yet functional)
cd apiGateway
go run cmd/main.go
```

## Testing with Postman

Since the API Gateway is not finished, you can directly test each service via gRPC:

## Get all products from InventoryService:
```
grpcurl -plaintext localhost:50051 inventory.InventoryService/ListProducts
```

## Create an order in OrderService:
```
grpcurl -plaintext -d '{
  "userId": 1,
  "items": [
    {"productId": 1, "quantity": 2}
  ]
}' localhost:50052 order.OrderService/CreateOrder
```

## Register a user in UserService:
```
grpcurl -plaintext -d '{
  "username": "admin",
  "password": "1234"
}' localhost:50053 user.UserService/Register
```

## Login with a user:
```
grpcurl -plaintext -d '{
  "username": "admin",
  "password": "1234"
}' localhost:50053 user.UserService/Login
```

## Notes
- All microservices use gRPC and follow Clean Architecture principles.

- API Gateway is still a work in progress and currently non-functional.

- Each service runs independently and listens on its own port.

- Authorization is handled via UserService (basic login flow).

- Extendable to include more features like payments, advanced auth, etc.
