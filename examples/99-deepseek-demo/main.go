package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

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

	fmt.Println("This is deepseek demo; type `bye` to exit")
	for {
		fmt.Println()
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		inputMessage := string(lineBytes)

		if strings.ToLower(inputMessage) == "bye" {
			os.Exit(0)
		}

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

		// call deepseek api
		sr, err := cli.StreamChatCompletionsReasoner(context.Background(), chatReq)
		if err != nil {
			panic(err)
		}

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
}
