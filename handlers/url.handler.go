package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sumitsj/url-shortener/contracts"
	"github.com/sumitsj/url-shortener/services"
	"net/http"
)

type urlHandler struct {
	service services.UrlService
}

func (receiver *urlHandler) ShortenUrl(ctx *gin.Context) {
	requestBody := contracts.ShortenUrlRequest{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errorMessage := "Failed to parse request body."
		fmt.Println(errorMessage)
		ctx.JSON(http.StatusBadRequest, contracts.ShortenUrlResponse{
			Error: errorMessage,
		})
		return
	}

	shortUrl := receiver.service.GenerateShortUrl(requestBody.URL)

	ctx.JSON(http.StatusOK, contracts.ShortenUrlResponse{
		OriginalUrl:  requestBody.URL,
		ShortenedUrl: shortUrl,
	})
}

type URLHandler interface {
	ShortenUrl(ctx *gin.Context)
}

func CreateUrlHandler(service services.UrlService) URLHandler {
	return &urlHandler{
		service: service,
	}
}
