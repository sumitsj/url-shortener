package repositories

import (
	"github.com/kamva/mgm/v3"
	"github.com/sumitsj/url-shortener/models"
)

type UrlMappingRepository interface {
	Create(urlMapping *models.URLMapping) error
}

type urlMappingRepository struct{}

func (r *urlMappingRepository) Create(urlMapping *models.URLMapping) error {
	err := mgm.Coll(urlMapping).Create(urlMapping)
	return err
}

func CreateUrlMappingRepository() UrlMappingRepository {
	return &urlMappingRepository{}
}
