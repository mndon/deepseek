package deepseek

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-deepseek/deepseek/internal"
)

const DEFAULT_TIMEOUT_SECONDS = 30

type client struct {
	*http.Client
	ApiKey string
}

func (c *client) CallChatCompletionsChat(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error) {
	// validate request
	if chatReq.Stream {
		return nil, errors.New(`err: stream should be "false"`)
	}
	if chatReq.Model != "deepseek-chat" {
		return nil, errors.New(`err: model should be "deepseek-chat"`)
	}
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}

	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	chatResp := &DeepseekChatResponse{}
	err = json.NewDecoder(respBody).Decode(chatResp)
	if err != nil {
		return nil, err
	}

	return chatResp, err
}

func (c *client) StreamChatCompletionsChat(chatReq *DeepseekChatRequest) (*MessageIterator, error) {
	if !chatReq.Stream {
		return nil, errors.New(`err: stream should be "true"`)
	}
	if chatReq.Model != "deepseek-chat" {
		return nil, errors.New(`err: model should be "deepseek-chat"`)
	}
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}

	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}

	msgIter := NewMessageIterator(respBody)
	return msgIter, nil
}

func (c *client) CallChatCompletionsReasoner(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error) {
	// validate request
	if chatReq.Stream {
		return nil, errors.New(`err: stream should be "false"`)
	}
	if chatReq.Model != "deepseek-reasoner" {
		return nil, errors.New(`err: model should be "deepseek-reasoner"`)
	}
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}

	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	chatResp := &DeepseekChatResponse{}
	err = json.NewDecoder(respBody).Decode(chatResp)
	if err != nil {
		return nil, err
	}

	return chatResp, err
}

func (c *client) StreamChatCompletionsReasoner(chatReq *DeepseekChatRequest) (*MessageIterator, error) {
	if !chatReq.Stream {
		return nil, errors.New(`err: stream should be "true"`)
	}
	if chatReq.Model != "deepseek-reasoner" {
		return nil, errors.New(`err: model should be "deepseek-reasoner"`)
	}
	err := ValidateRequest(chatReq)
	if err != nil {
		return nil, err
	}

	respBody, err := c.do(chatReq)
	if err != nil {
		return nil, err
	}

	msgIter := NewMessageIterator(respBody)
	return msgIter, nil
}

func (c *client) do(chatReq *DeepseekChatRequest) (io.ReadCloser, error) {
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
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errMsg))
	}

	return resp.Body, nil
}

func setDefaultHeaders(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}
