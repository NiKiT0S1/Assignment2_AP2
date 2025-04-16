// package cmd
package main

import (
	"log"
	"net"
	grpcDelivery "orderService/internal/delivery/grpc"
	pb "orderService/internal/delivery/grpc/pb"
	"orderService/internal/repository"
	"orderService/internal/usecase"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Подключение к БД
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln("DB connection failed:", err)
	}

	// Инициализация репозитория и бизнес-логики
	orderRepo := repository.NewOrderRepo(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)

	// Создание и запуск gRPC сервера
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Регистрация обработчика
	orderHandler := grpcDelivery.NewOrderHandler(orderUC)
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	log.Println("OrderService gRPC started on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
