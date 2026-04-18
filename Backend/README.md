# Ecommerce Backend

This is the Go backend for the ecommerce project. It exposes HTTP APIs for users and products, uses a simple JWT-based auth flow, and currently keeps data in memory rather than a persistent database.

## Features

- User registration and login
- JWT-protected routes for user and product management
- Product listing, details, create, update, and delete endpoints
- Built-in middleware for logging, CORS, and preflight handling
- Docker support for local containerized runs

## Requirements

- Go 1.25.5 or compatible
- A `.env` file in the `Backend/` directory

## Configuration

Create a `.env` file with these values:

```env
VERSION=1.0.0
SERVICE_NAME=ecommerce-backend
HTTP_PORT=3080
JWT_SECRET_KEY=your_secret_key
```

If any of these values are missing, the application exits on startup.

## Run Locally

From the `Backend/` directory:

```bash
go mod download
go run main.go
```

The server starts on the port set in `HTTP_PORT`.

## Run With Docker

From the `Backend/` directory:

```bash
docker build -t ecommerce-backend .
docker run --env-file .env -p 3080:3080 ecommerce-backend
```

## API Overview

### Users

- `POST /users` - register a new user
- `GET /users` - list users, requires JWT
- `POST /users/login` - authenticate and receive a JWT

Example login payload:

```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```

### Products

- `GET /products` - list products
- `GET /products/{productId}` - get a single product
- `POST /products` - create a product, requires JWT
- `PUT /products/{productId}` - update a product, requires JWT
- `DELETE /products/{productId}` - delete a product, requires JWT

Example product payload:

```json
{
  "title": "Apple",
  "description": "Fresh red apples",
  "price": 12.5,
  "imgUrl": "https://example.com/apple.jpg"
}
```

## Authentication

Protected routes expect the `Authorization` header in this format:

```http
Authorization: Bearer <jwt_token>
```

The token is signed with `JWT_SECRET_KEY` from the `.env` file.

## Data Storage

This version stores users and products in memory only. Data is reset when the server restarts. The product list starts with a few seeded sample items.

## Project Structure

```text
Backend/
  main.go              Application entry point
  cmd/                 Startup orchestration
  config/              Environment-based configuration loading
  database/            In-memory user and product storage
  rest/                HTTP server, handlers, and middleware
  util/                JWT and response helpers
```

## Notes

- The backend is designed to work with the frontend in the root workspace.
- If you want persistence later, the `database/` package is the place to replace with a real database layer.