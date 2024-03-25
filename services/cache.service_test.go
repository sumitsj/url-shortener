package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCacheService(t *testing.T) {
	service := NewCacheService()
	assert.NotNil(t, service)
}

func Test_cacheService_get_and_put(t *testing.T) {
	service := NewCacheService()

	err := service.put("key", "value")
	require.NoError(t, err)

	value, err := service.get("key")
	require.NoError(t, err)
	assert.Equal(t, "value", value)
}

func Test_cacheService_get_shouldThrowError(t *testing.T) {
	service := NewCacheService()

	_, err := service.get("key")
	assert.Equal(t, "key not found", err.Error())
}
