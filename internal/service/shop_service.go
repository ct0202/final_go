package service

import (
	"context"
	"log"

	"github.com/ct0202/bicycle-shop/internal/domain"
)

type ShopService struct {
	repo  Repository
	queue Queue
	email *EmailService
}

type Repository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
}

type Queue interface {
	PublishOrderNotification(ctx context.Context, orderID, userEmail string) error
}

func NewShopService(repo Repository, queue Queue) *ShopService {
	return &ShopService{
		repo:  repo,
		queue: queue,
	}
}

func (s *ShopService) CreateOrder(ctx context.Context, order *domain.Order) error {
	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return err
	}
	return nil
}

func (s *ShopService) PublishOrderNotification(ctx context.Context, orderID, userEmail string) error {
	if err := s.queue.PublishOrderNotification(ctx, orderID, userEmail); err != nil {
		log.Printf("Failed to publish notification: %v", err)
		return err
	}
	return nil
}

func (s *ShopService) ListBicycles(ctx context.Context) ([]*domain.Bicycle, error) {
	// TODO: Implement bicycle listing
	return []*domain.Bicycle{}, nil
}
