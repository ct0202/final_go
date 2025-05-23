# Bicycle Shop

## Project Overview
This is an online bicycle shop with rental functionality. Users can browse bicycles, place orders, and rent bikes. The system supports email notifications and is built with clean architecture principles.

## Technologies Used
- Go, gRPC, REST (optional)
- PostgreSQL, Redis
- RabbitMQ
- Gmail API for emails
- Grafana, Prometheus, OpenTelemetry (bonus)
- Docker

## How to Run Locally
1. Clone the repository: `git clone <repo-url>`
2. Set up environment variables (e.g., Gmail API credentials).
3. Run `docker-compose up --build`.
4. Access gRPC server on `localhost:8080`.

## How to Run Tests
```bash
go test ./tests -v