package repositories

import (
	"errors"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/sumitsj/url-shortener/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UrlMappingRepository interface {
	Create(urlMapping *models.URLMapping) error
	GetByShortKey(shortUrl string) (*models.URLMapping, error)
	GetByOriginalUrl(originalUrl string) (*models.URLMapping, error)
	GetTopNDomains(top int) []string
	GetAll() ([]models.URLMapping, error)
}

type urlMappingRepository struct{}

func (r *urlMappingRepository) GetAll() (list []models.URLMapping, err error) {
	err = mgm.Coll(&models.URLMapping{}).SimpleFind(
		&list,
		bson.M{},
		nil,
	)
	return
}

func (r *urlMappingRepository) GetTopNDomains(top int) []string {
	panic("not implemented")
}

func (r *urlMappingRepository) Create(urlMapping *models.URLMapping) error {
	err := mgm.Coll(urlMapping).Create(urlMapping)
	return err
}

func (r *urlMappingRepository) GetByOriginalUrl(originalUrl string) (*models.URLMapping, error) {
	urlMapping := &models.URLMapping{}

	if err := mgm.Coll(urlMapping).First(bson.M{"originalUrl": originalUrl}, urlMapping); err != nil {
		return nil, errors.New(fmt.Sprintf("GetByOriginalUrl: Can not find URL mapping for original URL: \"%v\". Internal Error: \"%v\"", originalUrl, err.Error()))
	}

	return urlMapping, nil
}

func (r *urlMappingRepository) GetByShortKey(shortUrl string) (*models.URLMapping, error) {
	urlMapping := &models.URLMapping{}

	if err := mgm.Coll(urlMapping).First(bson.M{"shortenedUrl": shortUrl}, urlMapping); err != nil {
		return nil, errors.New(fmt.Sprintf("GetByShortKey: Can not find URL mapping for short URL: \"%v\". Internal Error: \"%v\"", shortUrl, err.Error()))
	}

	return urlMapping, nil
}

func CreateUrlMappingRepository() UrlMappingRepository {
	return &urlMappingRepository{}
}
