package token

import (
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	requestHeaderClientID = "X-Client-Id"
)

var (
	// TokenHelper get and authenticate request
	tokenHelper accessToken
)

// AccessToken token for authentication
type AccessToken struct {
	TokenID string `json:"token_id"`
}

// AccessTokenInterface interface
type AccessTokenInterface interface {
	Authenticate(http.Request) error
	GetClientID(http.Request) string
}

type accessToken struct {
	restClient *rest.RequestBuilder
}

// GetTokenHelper Get token helper
func GetTokenHelper() AccessTokenInterface {
	return tokenHelper
}

func (at accessToken) Authenticate(r http.Request) error {
	response := at.restClient.Get("/abc")
	fmt.Println(response)
	return nil
}

func (at accessToken) GetClientID(r http.Request) string {
	return r.Header.Get(requestHeaderClientID)
}

func getAccessToken(token string) (*AccessToken, error) {
	tokenHelper.restClient.Get("/abc")
	return nil, nil
}
