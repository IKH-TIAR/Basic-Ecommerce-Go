# ecommerce

Lightweight e‑commerce backend written in Go. Provides user registration/login and product CRUD with simple JWT authentication and PostgreSQL persistence.

**Tech stack:** Go (module: `ecommerce`), PostgreSQL, `github.com/joho/godotenv`, `github.com/jmoiron/sqlx`, `github.com/lib/pq`.

**Go version:** 1.25.x (see `go.mod`).

**Quick start**

1. Create a `.env` file (example below).
2. Ensure PostgreSQL is running and accessible using the values in `.env`.
3. Build or run the server:

```bash
go build -o ecommerce ./...
./ecommerce
# or for development
go run main.go
```

The binary calls `cmd.Serve()` which initializes DB, repositories, services and starts the HTTP server.

**Environment variables (.env)**

Required variables (the app exits if any are missing):

- `VERSION` — application version
- `SERVICE_NAME` — service name
- `HTTP_PORT` — port the HTTP server will listen on (e.g. `8080`)
- `SECRET` — HMAC secret used to sign JWTs
- `DB_HOST` — Postgres host
- `DB_PORT` — Postgres port
- `DB_USER` — Postgres user
- `DB_PASSWORD` — Postgres password
- `DB_NAME` — Postgres database name
- `DB_SSL_MODE` — optional (defaults to `disable`)

Example `.env`:

```env
VERSION=0.1.0
SERVICE_NAME=ecommerce
HTTP_PORT=8080
SECRET=replace-with-secret
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce
DB_SSL_MODE=disable
```

**Database & migrations**

This project uses PostgreSQL. SQL migration files are in the `migrations/` folder and include:

- `000001_create_user_table.up.sql` — creates `users` table
- `000002_product_create_table.up.sql` — creates `products` table

Apply migrations with your preferred tool (e.g. `psql`, `migrate`, `golang-migrate`). The repository does not include an automatic migration runner.

**API endpoints**

Authentication: endpoints protected by JWT expect an `Authorization: Bearer <token>` header. Tokens are HMAC-SHA256 signed using the `SECRET` value.

User routes:

- POST /users — register a new user
- POST /users/login — login and receive a JWT

Product routes:

- GET /products — list products
- POST /products — create a product (requires Bearer token)
- GET /products/{id} — get product by id
- PUT /products/{id} — update product (requires Bearer token)
- DELETE /products/{id} — delete product (requires Bearer token)

Notes: routes are registered in `rest/handlers/*/routes.go`.

**Project structure (high level)**

- `main.go` — application entry (calls `cmd.Serve()`)
- `cmd/serve.go` — bootstraps config, DB, services and REST server
- `config/` — environment/config loader
- `infra/db/` — DB connection helper
- `repo/` — repository implementations for persistence
- `product/`, `user/` — domain services and ports
- `rest/` — HTTP server, handlers and middleware
- `utils/` — helpers (JWT creation, JSON/error writers)
- `migrations/` — SQL migration files

**Dependencies**

See `go.mod` for module and versions (uses `github.com/joho/godotenv`, `github.com/jmoiron/sqlx`, `github.com/lib/pq`).

**Design & Architecture**

- **Design pattern:** The project follows a simple layered architecture separating concerns across configuration, infrastructure, repository, domain (services/ports), and delivery (HTTP handlers). This keeps code testable and easier to refactor.
- **Repository pattern:** Persistence is encapsulated in the `repo/` package (e.g. `repo/user.go`, `repo/product.go`). Handlers and services depend on repository interfaces (ports) rather than concrete DB code, enabling easier swapping of DB implementations and simpler unit testing.
- **Domain-Driven Design (DDD) concepts:** The codebase applies light DDD principles:
	- Entities and domain logic live in `domain/` (`user.go`, `product.go`).
	- Application services live in `user/` and `product/` packages to coordinate use cases.
	- Ports/adapters are used: handlers and repo adapters implement ports to communicate with domain services.

