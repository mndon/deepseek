package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

var apiKey = "your_deepseek_api_key"

func main() {
	if apiKeyEnv := os.Getenv("DEEPSEEK_API_KEY"); apiKeyEnv != "" {
		apiKey = apiKeyEnv
	}

	// create deepseek api client
	cli, err := deepseek.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	inputMessage := "Hello" // set your input message

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_REASONER_MODEL,
		Stream: true,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
	}
	fmt.Printf("input message => %s\n", chatReq.Messages[0].Content)

	// call deepseek api
	sr, err := cli.StreamChatCompletionsReasoner(context.Background(), chatReq)
	if err != nil {
		panic(err)
	}

	fmt.Print("output message => ")
	for {
		chatResp, err := sr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if chatResp.Choices[0].Delta.ReasoningContent != "" {
			fmt.Print(chatResp.Choices[0].Delta.ReasoningContent)
		} else {
			fmt.Print(chatResp.Choices[0].Delta.Content)
		}
	}
}
