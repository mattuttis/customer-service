# Customer Service Microservice

## Overview

Production-ready Go microservice with REST API for creating and retrieving customers.

## Tech Stack

- **Go 1.25**
- **Gin** - Web framework
- **Prometheus** - Metrics
- **slog** - Structured logging
- **Docker** - Containerization
- **GitHub Actions** - CI/CD

## Architecture

Hexagonal (Ports & Adapters) / DDD:

```
cmd/api/main.go                          # Entry point, wiring
internal/
├── domain/                              # Core business logic
│   ├── customer.go                      # Entity + constructor + validation
│   └── repository.go                    # Port (interface)
├── application/                         # Use cases
│   └── customer_service.go              # Service orchestrating domain
├── adapter/
│   ├── inbound/handler/                 # REST handlers
│   │   ├── dto.go                       # Request/Response objects
│   │   ├── handler.go                   # HTTP handlers
│   │   └── router.go                    # Route setup
│   └── outbound/persistence/            # Repository implementation
│       └── memory.go                    # In-memory storage
└── config/
    └── config.go                        # Environment configuration
pkg/
├── logging/
│   ├── logger.go                        # slog JSON logger
│   └── middleware.go                    # Request logging middleware
└── metrics/prometheus/
    └── prometheus.go                    # Prometheus metrics + middleware
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | /customers | Create customer |
| GET | /customers/:id | Get customer by ID |
| GET | /health | Liveness probe |
| GET | /ready | Readiness probe |
| GET | /metrics | Prometheus metrics |

## Customer Entity

```go
type Customer struct {
    ID        uuid.UUID
    FirstName string    // required, 1-100 chars
    LastName  string    // required, 1-100 chars
    Age       int       // required, 18-150
}
```

## Running

```bash
# Local
make run

# Docker
make docker-build
make docker-run

# With custom port
SERVER_PORT=3000 make run
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| SERVER_PORT | 8080 | Server port |
| SHUTDOWN_TIMEOUT | 5s | Graceful shutdown timeout |

## Completed

- [x] Domain layer (entity, repository port)
- [x] Application layer (service)
- [x] Inbound adapter (REST handlers with Gin)
- [x] Outbound adapter (in-memory repository)
- [x] Input validation (binding tags)
- [x] Prometheus metrics
- [x] Health/readiness endpoints
- [x] Configuration from environment
- [x] Graceful shutdown
- [x] Structured logging (slog)
- [x] Dockerfile (multi-stage)
- [x] Makefile
- [x] CI/CD (GitHub Actions)

## TODO

- [ ] Unit tests (domain, application)
- [ ] Handler tests
- [ ] Integration tests

## Test Plan

1. **Domain tests** (`internal/domain/customer_test.go`)
   - Test NewCustomer with valid input
   - Test NewCustomer with invalid first name, last name, age

2. **Service tests** (`internal/application/customer_service_test.go`)
   - Test CreateCustomer success
   - Test CreateCustomer validation failure
   - Test GetCustomer found
   - Test GetCustomer not found
   - Use mock repository

3. **Handler tests** (`internal/adapter/inbound/handler/handler_test.go`)
   - Test Create endpoint (201, 400)
   - Test GetByID endpoint (200, 400, 404)
   - Use httptest

4. **Integration test**
   - Full flow: create customer, retrieve customer
