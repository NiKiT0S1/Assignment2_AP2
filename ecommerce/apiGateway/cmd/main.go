// package cmd
package main

import (
	"apiGateway/internal/delivery/handlers"
	"apiGateway/internal/grpc"
	"apiGateway/internal/middleware"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	// gRPC клиент UserService
	userClient, err := grpcDelivery.NewUserClient("localhost:50053")
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	// Middleware: авторизация через UserService
	r.Use(middleware.Auth(userClient))

	// Прокси
	inventoryURL, _ := url.Parse("http://localhost:8081")
	orderURL, _ := url.Parse("http://localhost:8082")

	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)

	r.Any("/products", gin.WrapH(inventoryProxy))
	r.Any("/products/*proxyPath", gin.WrapH(inventoryProxy))
	r.Any("/orders", gin.WrapH(orderProxy))
	r.Any("/orders/*proxyPath", gin.WrapH(orderProxy))

	// Маршруты для UserService
	handlers.RegisterRoutes(r, userClient)

	r.Run(":8080")
}
