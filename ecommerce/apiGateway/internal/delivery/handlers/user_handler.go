package handlers

import (
	grpcDelivery "apiGateway/internal/grpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userClient *grpcDelivery.UserClient) {
	r.POST("/register", func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		user, err := userClient.Register(body.Username, body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.Username})
	})

	r.POST("/login", func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		user, err := userClient.Authenticate(body.Username, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user.Id})
	})

	r.GET("/profile/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		user, err := userClient.GetProfile(int32(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": user.Id, "username": user.Username})
	})
}
