# Go Client - Deepseek API

### Why yet another Go client?

We needed to call the DeepSeek API from one of our Go services but couldn't find a complete and reliable Go client, so we built our own.

### Why this Go client is best?

Because this Go client is not only complete and reliable but also simple and performant.

- **Complete** – It offers full support for all APIs, including their complete request and response payloads. (Note: Beta version coming soon.)

- **Reliable** – We have implemented numerous Go tests to ensure that all features function correctly at all times.

- **Simple** – The client is organized into multiple Go packages, making it easy to navigate and ensuring that each package contains only relevant and necessary features.

- **Performant** – Speed is crucial when working with AI models. We have optimized this client to deliver the fastest possible performance.

## Install
```
go get github.com/go-deepseek/deepseek
```

## Usage

Here’s an example of sending a "Hello Deepseek!" message using `model=deepseek-chat` and `stream=false`

```
package main

import (
	"context"
	"fmt"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

func main() {
	cli, _ := deepseek.NewClient("your_deepseek_api_token")

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: "Hello Deepseek!", // set your input message
			},
		},
	}

	chatResp, err := cli.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		fmt.Println("Error =>", err)
		return
	}
	fmt.Printf("output => %s\n", chatResp.Choices[0].Message.Content)
}
```

## Examples

Please check the [examples](examples/) directory, which showcases each feature of this client.