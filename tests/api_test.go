//go:build integration

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sumitsj/url-shortener/contracts"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestUrlShortener(t *testing.T) {
	originalUrl := "www.google.com"
	baseUrl := "http://localhost:" + appPort
	url := baseUrl + "/short"

	log.Println("Calling: ", url)
	var body = []byte(fmt.Sprintf(`{"URL": "%s" }`, originalUrl))
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var response contracts.ShortenUrlResponse
	err = json.Unmarshal(bytes, &response)
	require.NoError(t, err)
	assert.NotEmpty(t, response)
	log.Println("URL Shortener Response : ", response)
}
