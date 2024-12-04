package utils

import (
	"encoding/json"
	"fmt"
	"github/mohanapranes/auth-filter-go/auth/config"
	"github/mohanapranes/auth-filter-go/auth/models"

	"net/http"
	"net/url"
	"strings"
)

type KeycloakClient struct {
	config *config.KeycloakConfig
}

func NewKeycloakClient(config *config.KeycloakConfig) *KeycloakClient {
	return &KeycloakClient{
		config: config,
	}
}

func (k *KeycloakClient) IntrospectToken(token string) (*models.TokenIntrospection, error) {
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", k.config.ClientID)
	data.Set("client_secret", k.config.ClientSecret)

	req, err := http.NewRequest("POST", k.config.IntrospectURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var tokenInfo models.TokenIntrospection
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &tokenInfo, nil
}
