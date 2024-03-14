package services

import (
	"fmt"
	"github.com/sumitsj/url-shortener/models"
	"github.com/sumitsj/url-shortener/repositories"
	"log"
	"math/rand"
	"time"
)

type UrlService interface {
	GenerateShortUrl(url string) string
}

type urlService struct {
	repository repositories.UrlMappingRepository
}

func (u *urlService) GenerateShortUrl(url string) string {
	shortKey := generateShortKey()

	shortenedURL := fmt.Sprintf("%v:%v/short/%s", Config.ServerAddr, Config.ServerPort, shortKey)

	urlMapping := models.URLMapping{
		OriginalUrl:  url,
		ShortenedUrl: shortenedURL,
	}

	if err := u.repository.Create(&urlMapping); err != nil {
		log.Printf("Failed to save url mapping for url: %v.\nError: %v", url, err)
	}

	return shortenedURL
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// TODO : Read key length from env if present
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func CreateUrlService(repository repositories.UrlMappingRepository) UrlService {
	return &urlService{
		repository: repository,
	}
}
