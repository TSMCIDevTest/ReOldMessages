package nina

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
)

type NinaAPI struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

type AuthResponse struct {
	Token string `json:"token"`
}

type MessageResponse struct {
	Status string `json:"status"`
}

func NewNinaAPI(baseURL, apiKey string) *NinaAPI {
	return &NinaAPI{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Authenticate sends an API key to the authentication endpoint and returns a token if successful.
func (n *NinaAPI) Authenticate() (string, error) {
	reqBody := map[string]string{"api_key": n.APIKey}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := n.HTTPClient.Post(n.BaseURL+"/auth", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.Token, nil
}

// SendMessage sends a message using the provided token and returns an error if the message could not be sent.
func (n *NinaAPI) SendMessage(token, message string) error {
	reqBody := map[string]string{"message": message}
	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", n.BaseURL+"/send", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var msgResp MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&msgResp); err != nil {
		return err
	}
	if msgResp.Status != "success" {
		return fmt.Errorf("failed to send message: %s", msgResp.Status)
	}

	return nil
}

// ReceiveMessages retrieves messages from the server using the provided token.
func (n *NinaAPI) ReceiveMessages(token string) ([]string, error) {
	req, err := http.NewRequest("GET", n.BaseURL+"/messages", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to receive messages: %s", resp.Status)
	}

	var messages []string
	if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}