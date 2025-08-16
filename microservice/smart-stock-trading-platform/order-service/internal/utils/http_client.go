package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetUserInfo(userServiceURL, userID, token string) (*UserResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/api/users/%s", userServiceURL, userID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("UserService error: %s", string(body))
		return nil, errors.New("unauthorized or failed response from user-service")
	}

	var userInfo UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
