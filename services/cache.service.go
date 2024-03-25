package services

import "errors"

type cacheService struct {
	store map[string]string
}

func (c *cacheService) get(key string) (string, error) {
	value, ok := c.store[key]

	if ok {
		return value, nil
	}

	return "", errors.New("key not found")
}

func (c *cacheService) put(key, value string) error {
	if c.store == nil {
		return errors.New("cache storage is not initialized")
	}

	c.store[key] = value
	return nil
}

type CacheService interface {
	get(key string) (string, error)
	put(key, value string) error
}

func NewCacheService() CacheService {
	return &cacheService{
		store: map[string]string{},
	}
}
