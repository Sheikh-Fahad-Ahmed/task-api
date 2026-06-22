# Task API

A lightweight RESTful API built using Gin Web Framework and SQLite3. It implements standard CRUD operations
for task management. It also features a custom middleware for logger designed to track exact
request latencies and write them straight into a local log file for each endpoint.

## Features

1. **Full CRUD Support**: Create, Read, Update, and Delete tasks dynamically.
2. **Persistent Storage**: Data is stored locally in an optimized SQLite database (`tasks.db`)
3. **Custom Logging Middleware**: Automatically captures and logs HTTP method, path, and execution duration for every incoming request.
4. **Clean Terminal Startup**: Disables cluttered terminal output; all operational HTTP telemetry is routed
   to local log file.

## Project Structure

```text
├── db/
│   └── db.go         # SQLite initialization & schema creation
├── handlers/
│   └── task_handlers.go   # HTTP request parsing, validation, and response handling
├── logger/
│   └── logger.go     # Custom Gin middleware for latency & error tracking
├── models/
│   └── task.go       # Struct data models for internal typing
├── store/
│   └── store.go      # SQL query executions and data layer operations
├── main.go           # Application entrypoint & route registry
├── gin.log           # Application runtime logs (Generated automatically, Git-ignored)
└── tasks.db          # SQLite Database file (Generated automatically, Git-ignored)
```

## Data Model

### Task Model

```Go
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

```

### Task Input Model

```Go
type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}
```

### Task Update Model

```Go
type TaskUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

```

### Task JSON Schema Model

```json
{
        "id": 1,
        "title": "task",
        "description": "one",
        "status": "in-progress",
        "created_at": "2026-06-22T08:34:45.065009+05:30"
}
```

## Installation & Setup

### Prerequisites

- **Go**: Ensure Go (v1.24 or higher) is installed.

### Steps to Run

1. **Clone the repository and navigate to the directory**

```bash
cd task-api
```

2. **Initialize modules and fetch dependencies**

```bash
go mod tidy
```

3. **Run the server**

```bash
go run main.go
```
