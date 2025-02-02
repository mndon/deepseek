package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/request"
)

var apiKey = "your_deepseek_api_key"

func main() {
	if apiKeyEnv := os.Getenv("DEEPSEEK_API_KEY"); apiKeyEnv != "" {
		apiKey = apiKeyEnv
	}

	// create deepseek api client with custom config
	cfg := config.Config{
		ApiKey:                   apiKey,
		TimeoutSeconds:           60,
		DisableRequestValidation: true,
	}
	cli, err := deepseek.NewClientWithConfig(cfg)
	if err != nil {
		panic(err)
	}

	inputMessage := "Hello" // set your input message

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
	}
	fmt.Printf("input message => %s\n", chatReq.Messages[0].Content)

	// call deepseek api
	chatResp, err := cli.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("output message => %s\n", chatResp.Choices[0].Message.Content)
}
