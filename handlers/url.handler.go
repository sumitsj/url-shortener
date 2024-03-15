package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sumitsj/url-shortener/constants"
	"github.com/sumitsj/url-shortener/contracts"
	"github.com/sumitsj/url-shortener/services"
	"log"
	"net/http"
)

type urlHandler struct {
	service services.UrlService
}

func (receiver *urlHandler) ShortenUrl(ctx *gin.Context) {
	requestBody := contracts.ShortenUrlRequest{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println(constants.RequestParsingErrorMessage)
		ctx.JSON(http.StatusBadRequest, contracts.ShortenUrlResponse{
			Error: constants.RequestParsingErrorMessage,
		})
		return
	}

	shortUrl, err := receiver.service.GenerateShortUrl(requestBody.URL)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, contracts.ShortenUrlResponse{
			Error: constants.InternalServerErrorMessage,
		})
	}

	ctx.JSON(http.StatusOK, contracts.ShortenUrlResponse{
		OriginalUrl:  requestBody.URL,
		ShortenedUrl: shortUrl,
	})
}

func (receiver *urlHandler) HandleRedirect(ctx *gin.Context) {
	shortKey := ctx.Param(constants.ShortKeyPathVariableName)
	shortUrl := receiver.service.FormatShortUrl(shortKey)

	url, err := receiver.service.GetOriginalUrlBy(shortUrl)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, contracts.ShortenUrlResponse{
			Error: constants.RedirectionErrorMessage,
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}

type URLHandler interface {
	ShortenUrl(ctx *gin.Context)
	HandleRedirect(ctx *gin.Context)
}

func CreateUrlHandler(service services.UrlService) URLHandler {
	return &urlHandler{
		service: service,
	}
}
