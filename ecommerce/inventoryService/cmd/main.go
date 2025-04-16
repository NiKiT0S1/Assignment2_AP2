// package cmd
package main

import (
	"log"
	"net"

	//_ "inventoryService/internal/delivery/grpc"
	grpcDelivery "inventoryService/internal/delivery/grpc"
	pb "inventoryService/internal/delivery/grpc/pb"
	"inventoryService/internal/repository"
	"inventoryService/internal/usecase"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect DB:", err)
	}

	productRepo := repository.NewProductRepo(db)
	productUC := usecase.NewProductUsecase(productRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpcDelivery.NewInventoryHandler(productUC)
	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, server)

	log.Println("InventoryService gRPC started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
