# pod-overlap-service

A Go service for checking if two date ranges overlap, built with [Gin](https://github.com/gin-gonic/gin) and following the Package Oriented Design principles inspired by [Ardan Labs](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html).

## Overview

This service exposes a RESTful API endpoint that accepts two date ranges and returns whether they overlap. It is designed for clarity, maintainability, and testability, using Go's package-oriented design philosophy.

## Package Oriented Design

This project is structured to group code by business domain and platform concerns, not by technical layer. This approach:
- Encourages encapsulation and separation of concerns
- Makes it easy to locate and reason about business logic
- Keeps platform-specific code (e.g., HTTP, database, logging) isolated from core logic
- Facilitates testing and future refactoring

### Project Structure

```
overlap-service/
├── cmd/overlap-service/         # Application entry point (main package)
├── internal/
│   ├── overlap/                 # Core domain logic (date range overlap)
│   └── platform/
│       ├── http/                # RESTful API, Gin setup, handlers, DTOs
│       └── logger/              # Logging implementation and configuration
├── go.mod, go.sum               # Go module files
├── Makefile                     # Build and run tasks
└── README.md                    # Project documentation
```

- **cmd/overlap-service/**: Contains the `main.go` file, which initializes the Gin engine and registers routes.
- **internal/overlap/**: Contains the core business logic for date range overlap, with no dependencies on platform code.
- **internal/platform/http/**: Contains HTTP handlers, DTOs, and Gin route registration. This is the only place that knows about the web framework.
- **internal/platform/logger/**: Contains your logger implementation and configuration.

## API

### POST /api/overlap

Checks if two date ranges overlap.

**Request JSON:**
```json
{
  "start_range": { "start": "2024-07-01T00:00:00Z", "end": "2024-07-10T00:00:00Z" },
  "end_range": { "start": "2024-07-05T00:00:00Z", "end": "2024-07-15T00:00:00Z" }
}
```

**Response JSON:**
```json
{
  "overlap": true
}
```

- Dates must be in RFC3339 format (e.g., `2024-07-01T00:00:00Z`).
- Returns 400 for invalid input or date order.

## Getting Started

1. **Install dependencies:**
   ```sh
   go mod tidy
   ```
2. **Run the service:**
   ```sh
   make run
   # or
   go run ./cmd/pod-overlap-service
   ```
3. **Test the API:**
   ```sh
   curl -X POST http://localhost:8080/api/overlap \
     -H 'Content-Type: application/json' \
     -d '{"start_range": {"start": "2024-07-01T00:00:00Z", "end": "2024-07-10T00:00:00Z"}, "end_range": {"start": "2024-07-05T00:00:00Z", "end": "2024-07-15T00:00:00Z"}}'
   # Response: {"overlap":true}
   ```
