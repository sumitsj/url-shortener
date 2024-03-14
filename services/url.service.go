package services

import (
	"fmt"
	"math/rand"
	"time"
)

type UrlService interface {
	GenerateShortUrl(url string) string
}

type urlService struct {
}

func (u *urlService) GenerateShortUrl(url string) string {
	shortKey := generateShortKey()

	// TODO : Read host from env if present
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// TODO : Save mapping of original & shortened URL

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

func CreateUrlService() UrlService {
	return &urlService{}
}
