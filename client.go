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

const DEFAULT_TIMEOUT_SECONDS = 30

type Client struct {
	http.Client
	ApiKey string
}

func NewClient(apiKey string) *Client {
	return NewClientWithTimeout(apiKey, DEFAULT_TIMEOUT_SECONDS)
}

func NewClientWithTimeout(apiKey string, timeoutSeconds int) *Client {
	timeout := time.Second * time.Duration(timeoutSeconds)
	c := &Client{
		ApiKey: apiKey,
		Client: http.Client{
			Timeout: timeout,
		},
	}
	return c
}

func (c *Client) Call(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error) {
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}
	if chatReq.Stream {
		return nil, errors.New(`err: stream should not be "true"`)
	}

	url := fmt.Sprintf(`%s/chat/completions`, BASE_URL)

	in := new(bytes.Buffer)
	err = json.NewEncoder(in).Encode(chatReq)
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

func (c *Client) Stream(chatReq *DeepseekChatRequest) (*MessageIterator, error) {
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}
	if !chatReq.Stream {
		return nil, errors.New(`err: stream should not be "false"`)
	}

	url := fmt.Sprintf(`%s/chat/completions`, BASE_URL)

	in := new(bytes.Buffer)
	err = json.NewEncoder(in).Encode(chatReq)
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

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errMsg))
	}

	msgIter := NewMessageIterator(resp.Body)
	return msgIter, nil
}

func SetDefaultHeaders(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}
