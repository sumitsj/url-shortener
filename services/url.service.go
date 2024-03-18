package services

import (
	"errors"
	"fmt"
	"github.com/sumitsj/url-shortener/models"
	"github.com/sumitsj/url-shortener/repositories"
	"log"
	"math/rand"
	"time"
)

type UrlService interface {
	GenerateShortUrl(url string) (string, error)
	GetOriginalUrlBy(shortKey string) (string, error)
	FormatShortUrl(shortKey string) string
}

type urlService struct {
	appConfig  AppConfig
	repository repositories.UrlMappingRepository
}

func (u *urlService) GenerateShortUrl(url string) (string, error) {
	urmMapping, err := u.repository.GetByOriginalUrl(url)

	if err != nil {
		log.Printf("Error while fetching URL mapping.\nError: %v", err.Error())

		shortKey := generateShortKey()

		urlMapping := models.URLMapping{
			OriginalUrl: url,
			ShortKey:    shortKey,
		}

		if err := u.repository.Create(&urlMapping); err != nil {
			errorMessage := fmt.Sprintf("Failed to save url mapping for url: %v.\nError: %v", url, err)
			return "", errors.New(errorMessage)
		}

		return u.FormatShortUrl(shortKey), nil
	}

	log.Printf("Url mapping found for original URL: \"%v\"", url)

	return u.FormatShortUrl(urmMapping.ShortKey), nil
}

func (u *urlService) GetOriginalUrlBy(shortKey string) (string, error) {
	urlMapping, err := u.repository.GetByShortKey(shortKey)
	if err != nil {
		return "", err
	}

	return urlMapping.OriginalUrl, nil
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
