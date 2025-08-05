# BackendEngineeringGo

A modern Go backend API service built with clean architecture principles, featuring user and post management with PostgreSQL integration.

## Features

- **RESTful API** with Chi router
- **PostgreSQL** database with connection pooling
- **Database migrations** using golang-migrate
- **Clean architecture** with separation of concerns
- **Docker** development environment
- **Environment-based** configuration
- **Comprehensive middleware** stack
- **Health check** endpoint

## Architecture

```mermaid
graph TB
    subgraph "Client Layer"
        C[HTTP Clients]
    end
    
    subgraph "API Layer (cmd/api)"
        R[Chi Router]
        M[Middleware Stack]
        H[HTTP Handlers]
        R --> M
        M --> H
    end
    
    subgraph "Business Layer (internal)"
        A[Application Logic]
        E[Environment Config]
        H --> A
        A --> E
    end
    
    subgraph "Data Layer (internal/store)"
        S[Storage Interface]
        US[UserStore]
        PS[PostStore]
        A --> S
        S --> US
        S --> PS
    end
    
    subgraph "Database Layer"
        DB[(PostgreSQL)]
        CONN[Connection Pool]
        US --> CONN
        PS --> CONN
        CONN --> DB
    end
    
    subgraph "Infrastructure"
        DC[Docker Compose]
        MIG[Migrations]
        DC --> DB
        MIG --> DB
    end
    
    C --> R
    
    style C fill:#e1f5fe
    style R fill:#f3e5f5
    style S fill:#e8f5e8
    style DB fill:#fff3e0
```

## Project Structure

```
BackendEngineeringGo/
├── cmd/                    # Application entry points
│   ├── api/               # Main API server
│   │   ├── main.go        # Application bootstrap
│   │   ├── api.go         # Server setup & routing
│   │   ├── health.go      # Health check handler
│   │   └── json.go        # JSON utilities
│   └── migrate/           # Database migrations
├── internal/              # Private application code
│   ├── db/               # Database connection logic
│   ├── env/              # Environment variable handling
│   └── store/            # Data access layer
│       ├── storage.go    # Storage interface
│       ├── users.go      # User model & operations
│       └── posts.go      # Post model & operations
├── web/                  # Web assets
├── scripts/              # Build/deployment scripts
├── docs/                 # Documentation
├── docker-compose.yaml   # PostgreSQL setup
├── Makefile             # Development commands
├── go.mod               # Go module definition
└── .env                 # Environment configuration
```

## Tech Stack

- **Language**: Go 1.24
- **Router**: Chi v5.2.2
- **Database**: PostgreSQL 16.3
- **Driver**: lib/pq v1.10.9
- **Config**: godotenv v1.5.1
- **Containerization**: Docker & Docker Compose
- **Migrations**: golang-migrate

## Getting Started

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- golang-migrate CLI tool

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/BackendEngineeringGo.git
   cd BackendEngineeringGo
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start PostgreSQL**
   ```bash
   docker-compose up -d
   ```

4. **Run database migrations**
   ```bash
   make migrate-up
   ```

5. **Install dependencies**
   ```bash
   go mod download
   ```

6. **Start the server**
   ```bash
   go run cmd/api/*.go
   ```

The API will be available at `http://localhost:8080`

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Posts Table
```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    title VARCHAR(255) NOT NULL,
    user_id INTEGER REFERENCES users(id),
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Development Commands

```bash
# Create a new migration
make migration create_users_table

# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down 1

# Run the application
go run cmd/api/*.go

# Run tests
go test ./...

# Build the application
go build -o bin/api cmd/api/*.go
```

## API Endpoints

### Health Check
```
GET /v1/health
```

**Response:**
```json
{
  "status": "ok",
  "environment": "development",
  "version": "0.0.1"
}
```

### Users (Coming Soon)
- `POST /v1/users` - Create user
- `GET /v1/users/{id}` - Get user
- `PUT /v1/users/{id}` - Update user
- `DELETE /v1/users/{id}` - Delete user

### Posts (Coming Soon)
- `POST /v1/posts` - Create post
- `GET /v1/posts` - List posts
- `GET /v1/posts/{id}` - Get post
- `PUT /v1/posts/{id}` - Update post
- `DELETE /v1/posts/{id}` - Delete post

## Configuration

Environment variables in `.env`:

```env
# Server Configuration
ADDR=:8080

# Database Configuration
DB_ADDR=postgres://admin:adminpassword@localhost:5432/backendengineeringGo?sslmode=disable
DB_MAX_OPEN_CONNS=30
DB_MAX_IDLE_CONNS=30
DB_MAX_IDLE_TIME=15m

# Application Environment
ENV=development
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Docker

The project includes a `docker-compose.yaml` for local development:

```bash
# Start PostgreSQL
docker-compose up -d

# Stop PostgreSQL
docker-compose down

# View logs
docker-compose logs -f
```

## Roadmap

- [ ] Complete CRUD operations for Users and Posts
- [ ] Authentication & Authorization (JWT)
- [ ] Input validation and sanitization
- [ ] Comprehensive error handling
- [ ] Unit and integration tests
- [ ] API documentation (OpenAPI/Swagger)
- [ ] Rate limiting and security middleware
- [ ] Logging and monitoring
- [ ] CI/CD pipeline
- [ ] Production Docker setup

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- Chi router for the excellent HTTP routing
- PostgreSQL team for the robust database
- Go community for the amazing ecosystem
