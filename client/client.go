package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/internal"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

type Client struct { // TODO: VN -- move to internal pkg
	*http.Client
	config.Config
}

func NewClient(config config.Config) (*Client, error) {
	if config.ApiKey == "" {
		return nil, errors.New("err: api key should not be blank")
	}
	if config.TimeoutSeconds == 0 {
		return nil, errors.New("err: timeout seconds should not be 0")
	}

	c := &Client{
		Config: config,
		Client: &http.Client{
			Timeout: time.Second * time.Duration(config.TimeoutSeconds),
		},
	}
	return c, nil
}

func (c *Client) CallChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	// validate request
	if chatReq.Stream {
		return nil, errors.New(`err: stream should be "false"`)
	}
	if chatReq.Model != "deepseek-chat" {
		return nil, errors.New(`err: model should be "deepseek-chat"`)
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	chatResp := &response.ChatCompletionsResponse{}
	err = json.NewDecoder(respBody).Decode(chatResp)
	if err != nil {
		return nil, err
	}

	return chatResp, err
}

func (c *Client) StreamChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	// validate request
	if !chatReq.Stream {
		return nil, errors.New(`err: stream should be "true"`)
	}
	if chatReq.Model != "deepseek-chat" {
		return nil, errors.New(`err: model should be "deepseek-chat"`)
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}

	sr := response.NewStreamReader(respBody)
	return sr, nil
}

func (c *Client) CallChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	// validate request
	if chatReq.Stream {
		return nil, errors.New(`err: stream should be "false"`)
	}
	if chatReq.Model != "deepseek-reasoner" {
		return nil, errors.New(`err: model should be "deepseek-reasoner"`)
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	chatResp := &response.ChatCompletionsResponse{}
	err = json.NewDecoder(respBody).Decode(chatResp)
	if err != nil {
		return nil, err
	}

	return chatResp, err
}

func (c *Client) StreamChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	// validate request
	if !chatReq.Stream {
		return nil, errors.New(`err: stream should be "true"`)
	}
	if chatReq.Model != "deepseek-reasoner" {
		return nil, errors.New(`err: model should be "deepseek-reasoner"`)
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}

	sr := response.NewStreamReader(respBody)
	return sr, nil
}

func (c *Client) do(chatReq *request.ChatCompletionsRequest) (io.ReadCloser, error) {
	url := fmt.Sprintf(`%s/chat/completions`, internal.BASE_URL)

	in := new(bytes.Buffer)
	err := json.NewEncoder(in).Encode(chatReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, in)
	if err != nil {
		return nil, err
	}
	setDefaultHeaders(req, c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return nil, processError(resp.Body, resp.StatusCode)
	}

	return resp.Body, nil
}

func setDefaultHeaders(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func processError(respBody io.Reader, statusCode int) error {
	errBody, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}
	errResp, err := internal.ParseError(errBody)
	if err != nil {
		return fmt.Errorf("err: %s; http_status_code=%d", errBody, statusCode)
	}
	return fmt.Errorf("err: %s; http_status_code=%d", errResp.Error.Message, statusCode)
}
