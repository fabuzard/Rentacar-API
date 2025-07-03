package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type VerifyRightResponse struct {
	Status bool `json:"status"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Checks struct {
		Format     bool `json:"format"`
		Domain     bool `json:"domain"`
		Disposable bool `json:"disposable"`
		SMTP       bool `json:"smtp"`
	} `json:"checks"`
}

func IsEmailValid(email string) (bool, string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	apikey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://verifyright.io/verify/%s?token=%s", email, apikey)

	resp, err := client.Get(url)
	if err != nil {
		return false, "", fmt.Errorf("failed to connect to email verification API: %w", err)
	}
	defer resp.Body.Close()

	var result VerifyRightResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, "", fmt.Errorf("failed to decode email verification response: %w", err)
	}

	if !result.Status {
		if result.Error.Message != "" {
			return false, result.Error.Message, nil
		}
		if !result.Checks.SMTP {
			return false, "email SMTP check failed", nil
		}
		return false, "email verification failed", nil
	}

	return true, "", nil
}
