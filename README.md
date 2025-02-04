# Go-Deepseek -- Go Client for [Deepseek API](https://api-docs.deepseek.com/)

> [!IMPORTANT]  
> We know that sometimes **Deepseek API is down** but we won't let you down.
>
> We have **`FakeCallbackClient`** using which you can continue your development and testing even though Deepseek API is down.
>
> Use `fake.NewFakeCallbackClient(fake.Callbacks{})` / See example [`examples/81-fake-callback-client/main_test.go`](examples/81-fake-callback-client/main_test.go)

## Demo

**30 seconds demo:** left-side browser with **[chat.deepseek.com](https://chat.deepseek.com/)** v/s **[go-deepseek](https://github.com/go-deepseek/deepseek)** in right-side terminal.

https://github.com/user-attachments/assets/baa05145-a13c-460d-91ce-90129c5b32d7

## Why yet another Go client?

We needed to call the DeepSeek API from one of our Go services but couldn't find a complete and reliable Go client, so we built our own.

## Why this Go client is better?

- **Complete** – It offers full support for all APIs, including their complete request and response payloads. (Note: Beta feature support coming soon.)

- **Reliable** – We have implemented numerous Go tests to ensure that all features work correctly at all times.

- **Simple** – The client is organized into multiple Go packages to ensure that each package contains only relevant and necessary features, making it easy to use.

- **Performant** – Speed is crucial when working with AI models. We have optimized this client to deliver the fastest possible performance.

## Install
```
go get github.com/go-deepseek/deepseek
```

## Usage

Here’s an example of sending a "Hello Deepseek!" message using `model=deepseek-chat` (**DeepSeek-V3 model**) and `stream=false`

```
package main

import (
	"context"
	"fmt"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

func main() {
	client, _ := deepseek.NewClient("your_deepseek_api_token")

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

	chatResp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		fmt.Println("Error =>", err)
		return
	}
	fmt.Printf("output => %s\n", chatResp.Choices[0].Message.Content)
}
```

Try above example:
```
First, copy above code in `main.go`
Replace `your_deepseek_api_token` with valid api token

$ go mod init
$ go get github.com/go-deepseek/deepseek

$ go run main.go
output => Hello! How can I assist you today? 😊
```

## Examples

Please check the [examples](examples/) directory, which showcases each feature of this client.

![examples](https://github.com/user-attachments/assets/032ff864-7da5-4b76-9484-836b52046614)

## Buy me a GitHub Star ⭐

If you like our work then please give github star to this repo. 😊
