package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robot-app/dto"
	"robot-app/pkg/config"
)

func ApiKeyAuthentication(c *gin.Context) {
	secretKey := c.GetHeader("x-api-key")
	if secretKey != config.GetAppConfig().ApiKey {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Message: "Unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
