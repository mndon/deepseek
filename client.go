package deepseek

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	http.Client

	ApiKey string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.Client = http.Client{
		Timeout: time.Second * 10,
	}
	c.ApiKey = apiKey
	return c
}

func (c *Client) Call(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error) {
	url := fmt.Sprintf(`%s/chat/completions`, BASE_URL)

	in := new(bytes.Buffer)
	err := json.NewEncoder(in).Encode(chatReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, in)
	if err != nil {
		return nil, err
	}
	SetDefaultHeaders(req, c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errMsg))
	}

	chatResp := &DeepseekChatResponse{}
	err = json.NewDecoder(resp.Body).Decode(chatResp)
	if err != nil {
		return nil, err
	}

	return chatResp, err
}

func SetDefaultHeaders(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}
