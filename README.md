# Task API

A lightweight RESTful API built using Gorilla/mux package. It implements standard CRUD operations
for task management. It utilizes an in-memory store, also features a custom middleware for logging
each endpoint.

## Features

1. **Full CRUD Support**: Create, Read, Update, and Delete tasks dynamically.
2. **In-Memory Store**: Fast performance using native Go slices (`[]models.Task`).
3. **Custom Logging Middleware**: Automatically captures and logs HTTP method, path, and execution duration for every incoming request.

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
