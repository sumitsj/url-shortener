package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitsj/url-shortener/handlers"
	"github.com/sumitsj/url-shortener/repositories"
	"github.com/sumitsj/url-shortener/router"
	"github.com/sumitsj/url-shortener/services"
)

func Start() error {
	services.LoadConfig()
	services.InitMongoDB()
	repository := repositories.CreateUrlMappingRepository()
	service := services.CreateUrlService(repository)
	handler := handlers.CreateUrlHandler(service)
	r := router.CreateRouter(handler).Init()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run()
}
