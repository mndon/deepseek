package main

import (
	"context"
	"testing"

	"github.com/go-deepseek/deepseek/fake"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

func TestGreeter(t *testing.T) {
	callbacks := fake.Callbacks{}

	callbacks.CallChatCompletionsChatCallback = func(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
		if chatReq.Messages[0].Content == "Hello" {
			chatResp := &response.ChatCompletionsResponse{
				Choices: []*response.Choice{
					{
						Message: &response.Message{
							Content: "How are you?",
						},
					},
				},
			}
			return chatResp, nil
		}

		if chatReq.Messages[0].Content == "Bye" {
			chatResp := &response.ChatCompletionsResponse{
				Choices: []*response.Choice{
					{
						Message: &response.Message{
							Content: "Good Day!",
						},
					},
				},
			}
			return chatResp, nil
		}

		return nil, nil
	}

	client := fake.NewFakeCallbackClient(callbacks)

	reply := Greeter(client, "Hello")
	if reply != "How are you?" {
		t.Fail()
	}

	reply = Greeter(client, "Bye")
	if reply != "Good Day!" {
		t.Fail()
	}
}
