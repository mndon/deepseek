package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mndon/deepseek"
	"github.com/mndon/deepseek/request"
)

func main() {
	// create deepseek api client
	cli, _ := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))

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
			fmt.Println("error => ", err)
			return
		}

		for {
			chatResp, err := sr.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("error => ", err)
				return
			}
			if chatResp.Choices[0].Delta.ReasoningContent != "" {
				fmt.Print(chatResp.Choices[0].Delta.ReasoningContent)
			} else {
				fmt.Print(chatResp.Choices[0].Delta.Content)
			}
		}
	}
}
