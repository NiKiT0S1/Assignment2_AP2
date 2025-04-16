package middleware

import (
	_ "apiGateway/internal/grpc"
	"apiGateway/internal/proto"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем Authorization заголовок
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Декодируем базовый авторизационный заголовок
		payload, _ := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Проверка через gRPC
		req := &proto.AuthRequest{
			Username: pair[0],
			Password: pair[1],
		}

		// Запрашиваем у UserService авторизацию
		resp, err := userClient.Authenticate(c, req)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Если аутентификация успешна, извлекаем user_id и передаем его в контекст
		c.Set("user_id", strconv.Itoa(int(resp.GetId())))

		// Переходим к следующему обработчику
		c.Next()
	}
}
