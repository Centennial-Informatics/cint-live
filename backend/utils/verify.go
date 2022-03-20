package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var httpClient = &http.Client{}

func VerifyToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.NewService(context.Background(), option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()

	tokenInfoCall.IdToken(idToken)

	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
