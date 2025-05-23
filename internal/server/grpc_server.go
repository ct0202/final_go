package server

import (
	"context"
	"log"
	"net"

	"github.com/ct0202/bicycle-shop/api/proto"
	"github.com/ct0202/bicycle-shop/internal/domain"
	"github.com/ct0202/bicycle-shop/internal/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	shopService  *service.ShopService
	emailService *service.EmailService
	proto.UnimplementedShopServiceServer
}

func NewGRPCServer(shopService *service.ShopService, emailService *service.EmailService) *GRPCServer {
	return &GRPCServer{
		shopService:  shopService,
		emailService: emailService,
	}
}

func (s *GRPCServer) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	proto.RegisterShopServiceServer(grpcServer, s)
	log.Printf("gRPC server running on %s", addr)
	return grpcServer.Serve(lis)
}

// Реализация gRPC-эндпоинтов
func (s *GRPCServer) ListBicycles(ctx context.Context, req *proto.ListBicyclesRequest) (*proto.ListBicyclesResponse, error) {
	bicycles, err := s.shopService.ListBicycles(ctx)
	if err != nil {
		return nil, err
	}

	resp := &proto.ListBicyclesResponse{}
	for _, b := range bicycles {
		resp.Bicycles = append(resp.Bicycles, &proto.Bicycle{
			Id:        b.ID,
			Name:      b.Name,
			Type:      b.Type,
			Price:     b.Price,
			Available: b.Available,
		})
	}
	return resp, nil
}

func (s *GRPCServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.Order, error) {
	order := &domain.Order{
		ID:         req.Id,
		UserID:     req.UserId,
		BicycleID:  req.BicycleId,
		TotalPrice: req.TotalPrice,
		Status:     "created",
	}

	if err := s.shopService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	// Асинхронно отправляем уведомление в очередь
	if err := s.shopService.PublishOrderNotification(ctx, order.ID, req.UserEmail); err != nil {
		log.Printf("Failed to publish notification: %v", err)
	}

	return &proto.Order{
		Id:         order.ID,
		UserId:     order.UserID,
		BicycleId:  order.BicycleID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, nil
}

func (s *GRPCServer) SendOrderConfirmation(ctx context.Context, req *proto.SendOrderConfirmationRequest) (*proto.Empty, error) {
	if err := s.emailService.SendOrderConfirmation(ctx, req.Email, req.OrderId); err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

// Другие методы (GetBicycle, CreateRental, etc.) опущены для краткости
// Они реализуются аналогично, вызывая методы shopService и конвертируя данные в proto-структуры

func (s *GRPCServer) Close() {
	// Закрытие соединений, если нужно
}
