package nina

import (
	"bytes"
	"encoding/json"
	"net/http"
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
		HTTPClient: &http.Client{},
	}
}

func (n *NinaAPI) Authenticate() (string, error) {
	reqBody := map[string]string{"api_key": n.APIKey}
	body, _ := json.Marshal(reqBody)

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
		return fmt.Errorf("failed to send message")
	}

	return nil
}

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
	defer resp.Body.Close()

	var messages []string
	if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}