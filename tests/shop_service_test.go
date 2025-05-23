package tests

import (
	"context"
	"testing"

	"github.com/ct0202/bicycle-shop/internal/domain"
	"github.com/ct0202/bicycle-shop/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestShopService_CreateOrder(t *testing.T) {
	svc := service.NewShopService(nil, nil, nil) // Моки для репозитория и очереди
	ctx := context.Background()

	order := &domain.Order{
		ID:         "order-1",
		UserID:     "user-1",
		BicycleID:  "bike-1",
		TotalPrice: 999.99,
		Status:     "created",
	}

	err := svc.CreateOrder(ctx, order)
	assert.NoError(t, err)
}
