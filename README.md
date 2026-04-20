# Task Manager API

A RESTful API for task management built with **Go** and **Clean Architecture** principles.

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Containerization:** Docker
- **Testing:** Testify + Mockery (75% coverage)

## Architecture

This project follows Clean Architecture with clear separation of concerns:

```
internal/
├── domain/       # Entities & interfaces
├── repository/   # Database layer
├── usecase/      # Business logic
└── handler/      # HTTP layer
```

Request flow:
```
HTTP Request → Handler → Usecase → Repository → Database
```

## Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose

### Installation

1. Clone the repository
```bash
git clone https://github.com/gummymule/task-manager.git
cd task-manager
```

2. Copy environment file
```bash
cp .env.example .env
```

3. Start PostgreSQL
```bash
docker-compose up -d
```

4. Run database migration
```bash
docker exec -i taskmanager_db psql -U postgres -d taskmanager < db/migrations/init.sql
```

5. Start the server
```bash
go run cmd/main.go
```

Server will run on `http://localhost:8080`

## API Endpoints

### Auth
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/register` | Register new user |
| POST | `/api/v1/login` | Login & get JWT token |

### Tasks (Protected — requires Bearer token)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/tasks` | Get all tasks (with pagination) |
| GET | `/api/v1/tasks/:id` | Get task by ID |
| POST | `/api/v1/tasks` | Create new task |
| PUT | `/api/v1/tasks/:id` | Update task |
| DELETE | `/api/v1/tasks/:id` | Delete task |

### Pagination
```
GET /api/v1/tasks?page=1&limit=10
```

## Request & Response Examples

### Register
```json
POST /api/v1/register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Login
```json
POST /api/v1/login
{
  "email": "john@example.com",
  "password": "password123"
}

Response:
{
  "status": "success",
  "message": "login success",
  "data": {
    "token": "eyJhbGci..."
  }
}
```

### Create Task
```json
POST /api/v1/tasks
Authorization: Bearer <token>

{
  "title": "Learn Clean Architecture",
  "description": "Implement clean architecture with Go",
  "status": "to_do"
}
```

### Task Status Values
| Status | Description |
|--------|-------------|
| `to_do` | Task not started |
| `in_progress` | Task in progress |
| `done` | Task completed |

## Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./internal/usecase/... -cover

# Run with verbose output
go test ./internal/usecase/... -v
```

## Project Structure

```
task-manager/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── db/
│   └── migrations/
│       └── init.sql
├── internal/
│   ├── domain/
│   │   ├── task.go
│   │   └── user.go
│   ├── handler/
│   │   ├── task_handler.go
│   │   └── user_handler.go
│   ├── mocks/
│   ├── repository/
│   │   ├── task_repository.go
│   │   └── user_repository.go
│   └── usecase/
│       ├── task_usecase.go
│       ├── task_usecase_test.go
│       ├── user_usecase.go
│       └── user_usecase_test.go
├── pkg/
│   ├── middleware/
│   │   └── auth.go
│   └── response/
│       └── response.go
├── .env.example
├── docker-compose.yml
├── go.mod
└── README.md
```