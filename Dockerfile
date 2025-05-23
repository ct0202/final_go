FROM golang:1.24.0

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bicycle-shop ./cmd/server

EXPOSE 8080
CMD ["/bicycle-shop"]