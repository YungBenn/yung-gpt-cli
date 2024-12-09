package gpt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Result string `json:"result"`
}

func SendMessage(message string) (string, error) {
	url := "https://chatgpt-42.p.rapidapi.com/gpt4"
	payload := strings.NewReader(fmt.Sprintf("{\"messages\":[{\"role\":\"user\",\"content\":\"%s\"}],\"web_access\":false}", message))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", "chatgpt-42.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Result, nil
}
