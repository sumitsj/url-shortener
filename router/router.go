package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitsj/url-shortener/handlers"
)

type Router struct {
	handler handlers.URLHandler
}

func (router *Router) Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/short", router.handler.ShortenUrl)
	return r
}

func CreateRouter(handler handlers.URLHandler) *Router {
	return &Router{
		handler: handler,
	}
}
