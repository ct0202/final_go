package main

import (
	"log"
	"os"

	"github.com/ct0202/bicycle-shop/internal/queue"
	"github.com/ct0202/bicycle-shop/internal/repository"
	"github.com/ct0202/bicycle-shop/internal/server"
	"github.com/ct0202/bicycle-shop/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// ctx := context.Background()

	// Подключение к PostgreSQL
	dbConnStr := "host=" + os.Getenv("DB_HOST") + " user=user password=password dbname=shop sslmode=disable"
	postgresRepo, err := repository.NewPostgresRepository(dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresRepo.Close()

	// Подключение к Redis
	redisRepo, err := repository.NewRedisRepository(os.Getenv("REDIS_HOST") + ":6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisRepo.Close()

	// Подключение к RabbitMQ
	rabbitMQ, err := queue.NewRabbitMQ(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Настройка Gmail API
	// emailService, err := service.NewEmailService(ctx, os.Getenv("GOOGLE_CREDENTIALS"))
	// if err != nil {
	// 	log.Fatalf("Failed to initialize email service: %v", err)
	// }

	// Инициализация сервиса магазина
	shopService := service.NewShopService(postgresRepo, rabbitMQ)

	// Запуск gRPC-сервера
	grpcServer := server.NewGRPCServer(shopService)
	if err := grpcServer.Run(":8080"); err != nil {
		log.Fatalf("Failed to run gRPC server: %v", err)
	}
}
