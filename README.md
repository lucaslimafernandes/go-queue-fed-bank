# go-queue-fed-bank

Simulation of the Caixa Econômica Federal queue system

Author: Lucas Lima Fernandes - lucas.lfernandes@live.com

## Project Structure

```plaintext
queue-simulator/
│
├── api/                 # Directory for the API
│   ├── main.go          # Main file to start the server
│   ├── config/          # Server configurations and other parameters
│   │   └── config.go
│   ├── handlers/        # Request handlers for the endpoints
│   │   ├── enter.go
│   │   ├── check.go
│   │   ├── use.go
│   │   └── release.go
│   ├── middleware/      # Middleware for authentication, logging, etc.
│   │   └── auth.go
│   └── model/           # Data models
│       └── queue.go
│
├── app/                 # Client application that simulates the requests
│   ├── main.go          # Entry point for the client application
│   ├── config/          # Client configurations, such as API URL
│   │   └── config.go
│   └── simulation/      # Logic for simulating the requests
│       ├── enter.go
│       ├── check.go
│       ├── use.go
│       └── release.go
│
├── pkg/                 # Packages shared between the API and the app
│   └── common/          # Common functions, like HTTP utilities
│       └── http.go
│
└── docker-compose.yml   # Docker Compose to orchestrate the API and possible databases

```

## Dev

docker run -d --name meu-redis-container -p 6379:6379 --rm -e REDIS_PASSWORD=suasenha redis:7 redis-server --requirepass suasenha

export PORT=":5000"
export REDIS_PASSWORD=suasenha

