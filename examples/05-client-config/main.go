package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mndon/deepseek"
	"github.com/mndon/deepseek/config"
	"github.com/mndon/deepseek/request"
)

func main() {
	// create deepseek api client with custom config
	cfg := config.Config{
		ApiKey:                   os.Getenv("DEEPSEEK_API_KEY"),
		TimeoutSeconds:           60,
		DisableRequestValidation: true,
	}
	cli, _ := deepseek.NewClientWithConfig(cfg)

	inputMessage := "Hello Deepseek!" // set your input message
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
		fmt.Println("error => ", err)
		return
	}
	fmt.Printf("output message => %s\n", chatResp.Choices[0].Message.Content)
}
