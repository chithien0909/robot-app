package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"robot-app/internal/http/handler"
	"robot-app/internal/http/middleware"
	"robot-app/pkg/config"
)

type Server interface {
	Start() error
}

type serverImpl struct {
	router  *gin.Engine
	handler handler.Handler
}

func New(h handler.Handler) Server {
	s := &serverImpl{
		handler: h,
	}
	s.withRouter()
	return s
}

func (i *serverImpl) Start() error {
	log.Printf("Listening and serving HTTP on :%d\n", config.GetAppConfig().Port)

	return i.router.Run(fmt.Sprintf(":%d", config.GetAppConfig().Port))
}

func (i *serverImpl) withRouter() {
	if config.GetAppConfig().Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	api := router.Group("/api", middleware.ApiKeyAuthentication)

	api.GET("/devices", i.handler.DeviceHandler().Find)

	i.router = router
}
