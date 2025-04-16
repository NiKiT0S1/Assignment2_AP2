package grpcDelivery

import (
	"apiGateway/internal/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

type UserClient struct {
	client proto.UserServiceClient
}

func NewUserClient(address string) (*UserClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // В реальной ситуации использовать grpc.WithTransportCredentials
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
		return nil, err
	}

	client := proto.NewUserServiceClient(conn)
	return &UserClient{client}, nil
}

// Обновляем метод Authenticate для работы с context и AuthRequest
func (u *UserClient) Authenticate(username, password string) (*proto.UserResponse, error) {
	req := &proto.AuthRequest{
		Username: username,
		Password: password,
	}

	// Контекст для gRPC запроса
	ctx := context.Background()

	// Вызов Authenticate с правильными параметрами
	return u.client.Authenticate(ctx, req)
}

// Add this method to fully implement the UserServiceClient interface
func (u *UserClient) Register(ctx context.Context, req *proto.RegisterRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return u.client.Register(ctx, req, opts...)
}

// Add this method
func (u *UserClient) GetProfile(ctx context.Context, req *proto.UserID, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return u.client.GetProfile(ctx, req, opts...)
}
