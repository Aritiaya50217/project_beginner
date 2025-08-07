package remote

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthClient interface {
	IsTokenValid(ctx *gin.Context, token string) (bool, error)
}

type authClient struct {
	userServiceURL string
}

type TokenValidationResponse struct {
	Valid bool `json:"valid"`
}

func NewAuthClient(userServiceURL string) AuthClient {
	return &authClient{
		userServiceURL: userServiceURL,
	}
}

func (a *authClient) IsTokenValid(ctx *gin.Context, token string) (bool, error) {
	req, err := http.NewRequest("POST", a.userServiceURL+"/v1/api/auth/validate", bytes.NewBuffer(nil))
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("token is invalid or expired")
	}

	var result TokenValidationResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.Valid, nil
}
