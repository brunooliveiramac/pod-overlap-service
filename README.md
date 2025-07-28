# pod-overlap-service

A Go service for checking if two date ranges overlap, following the Package Oriented Design principles inspired by [Ardan Labs](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html).
Go is a multiparadigm language that supports various programming styles, object-oriented, and functional programming. 
This service leverages Go's strengths to create a clear and maintainable codebase. You may see concepts from SOLID principles, such as Single Responsibility Principle (SRP), Dependency Inversion Principle (DIP), and Interface Segregation Principle (ISP) being
applied throughout the codebase.

## Overview

This service exposes a RESTful API endpoint that accepts two date ranges and returns whether they overlap.

## Package Oriented Design

This project is structured to group code by business domain and platform concerns, not by technical layer. This approach:
- Focus on the business domain rather than technical implementation details
- Encourages encapsulation and separation of concerns
- Makes it easy to locate and reason about business logic, screaming architecture concept
- Keeps platform-specific code (HTTP, database) isolated from core logic

### Project Structure

```
pod-overlap-service/
├── cmd/pod-overlap-service/     # Application entry point, can be HTTP server, CLI, etc. This couples with internal core
├── internal/
│   ├── overlap/                 # Core domain logic, uses interfaces for dependencies, to allow easy testing and mocking,and to keep it platform-agnostic, using dependency inversion principle (DIP)      
│   │   ├── overlap.go           # Core logic for checking date range overlap, it must follow Single Responsibility Principle (SRP) and Inteface Segregation Principle (ISP) to coupling with only what is needed.                    
│   └── platform/                # Platform-specific code, grouped by functionality, implements interfaces defined in core logic (DIP)   
│       ├── http/                #  - RESTful API,handlers, DTOs
│       │   ├── e2e_test.go      #  - End-to-end tests for the HTTP API, should be small and test quite few happy paths
│       └── logger/              #  - Logging implementation and configuration
│       └── dataprovider         #  - Data provider implementation (e.g., database, external APIs) also implements interfaces defined in core logic, so it can be easily mocked in tests keeping the core clean and maintainable from platform concerns.
├── Makefile                     # Build and run tasks
```

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
- Returns 200 with `{"overlap": true}` or `{"overlap": false}`.

## Getting Started

1. **Install dependencies:**
   ```sh
   go mod tidy
   ```
2. **Run the service:**
   ```sh
   make run
   ```
3. **Test the API:**
   ```sh
   curl -X POST http://localhost:8080/api/overlap \
     -H 'Content-Type: application/json' \
     -d '{"start_range": {"start": "2024-07-01T00:00:00Z", "end": "2024-07-10T00:00:00Z"}, "end_range": {"start": "2024-07-05T00:00:00Z", "end": "2024-07-15T00:00:00Z"}}'
   # Response: {"overlap":true}
   ```

## Testing

Run all tests across the project:

```sh
make test
```
