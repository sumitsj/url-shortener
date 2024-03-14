package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitsj/url-shortener/handlers"
	"github.com/sumitsj/url-shortener/router"
	"github.com/sumitsj/url-shortener/services"
)

func Start() error {
	service := services.CreateUrlService()
	handler := handlers.CreateUrlHandler(service)
	r := router.CreateRouter(handler).Init()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run()
}
