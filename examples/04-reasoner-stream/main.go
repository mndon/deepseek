package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/mndon/deepseek"
	"github.com/mndon/deepseek/request"
)

func main() {
	// create deepseek api client
	cli, _ := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))

	inputMessage := "Hello Deepseek!" // set your input message
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
		fmt.Println("error => ", err)
		return
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
