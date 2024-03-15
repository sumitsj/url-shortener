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
	GetOriginalUrlBy(shortenedUrl string) (string, error)
	FormatShortUrl(shortKey string) string
}

type urlService struct {
	appConfig  AppConfig
	repository repositories.UrlMappingRepository
}

func (u *urlService) GenerateShortUrl(url string) string {
	urmMapping, err := u.repository.GetByOriginalUrl(url)

	if err != nil {
		shortKey := generateShortKey()

		shortenedURL := u.FormatShortUrl(shortKey)

		urlMapping := models.URLMapping{
			OriginalUrl:  url,
			ShortenedUrl: shortenedURL,
		}

		if err := u.repository.Create(&urlMapping); err != nil {
			log.Printf("Failed to save url mapping for url: %v.\nError: %v", url, err)
		}

		return shortenedURL
	}

	log.Printf("Url mapping found for original URL: \"%v\"", url)

	return urmMapping.ShortenedUrl
}

func (u *urlService) GetOriginalUrlBy(shortenedUrl string) (string, error) {
	url, err := u.repository.GetByShortenedUrl(shortenedUrl)
	if err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}

func (u *urlService) FormatShortUrl(shortKey string) string {
	shortenedURL := fmt.Sprintf("%v:%v/s/%s", u.appConfig.GetServerAddr(), u.appConfig.GetServerPort(), shortKey)
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

func CreateUrlService(appConfig AppConfig, repository repositories.UrlMappingRepository) UrlService {
	return &urlService{
		appConfig:  appConfig,
		repository: repository,
	}
}
