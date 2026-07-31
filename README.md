# URLshortener

> A Go-based URL shortener service built with Chi router.

## 📑 Table of Contents

- [Key Features](#key-features)
- [Use Cases](#use-cases)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [License](#license)

## ✨ Key Features

- **🔗 Chi HTTP Router** — Handles HTTP routing and middleware components using the Chi v5 router library.
- **📁 Configurable Storage Backend** — Initializes storage based on file path configurations passed during application startup.
- **📝 Structured System Logging** — Logs application events with configurable log levels and dedicated file outputs.
- **🐳 Docker Compose Ready** — Includes Docker and Docker Compose configurations for containerized deployment.

## 🎯 Use Cases

- Running a self-hosted URL shortening API service within a containerized environment.
- Using as a reference template for building Go web APIs with Chi and structured storage.

## 🛠️ Tech Stack

- **Docker**
- **Go(Chi)**
- **Postgres**


## 📁 Project Structure

```
.
├── LICENSE
├── docker-compose.yaml
└── src
    ├── Dockerfile
    ├── cmd
    │   └── main.go
    ├── config
    │   ├── config.go
    │   └── config.yaml
    ├── go.mod
    ├── go.sum
    └── internal
        ├── server
        │   ├── handlers
        │   │   ├── alias_handlers.go
        │   │   ├── redirect_handlers.go
        │   │   └── sender.go
        │   └── mw_components
        │       └── logger.go
        └── storage
            ├── delete_url.go
            ├── entity.go
            ├── get_url.go
            ├── save_url.go
            └── setup_storage.go
```

## 📜 License

This project is licensed under the **GPL** License.
