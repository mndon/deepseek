package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

/*
Example Details:

Example shows how to use FakeCallbackClient to test user feature which is using go-deepseek client.

Here, Greeter() is user feature when invoke from main() needs deepseek client with DEEPSEEK_API_KEY

Check TestGreeter() in main_test.go file.  It is testing user feature Greeter() using FakeCallbackClient.  This does not need deepseek with DEEPSEEK_API_KEY.

Using FakeCallbackClient, you will be able to develop your feature and test your feature even though deepseek API is down.
*/

func main() {
	client, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		panic(err)
	}
	reply := Greeter(client, "Hello")
	fmt.Println(reply)
}

func Greeter(client deepseek.Client, message string) string {
	chatReq := &request.ChatCompletionsRequest{
		Model: deepseek.DEEPSEEK_CHAT_MODEL,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: message,
			},
		},
		Stream: false,
	}

	resp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		panic(err)
	}

	reply := resp.Choices[0].Message.Content
	return reply
}
