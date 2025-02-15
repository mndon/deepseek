# Go-Deepseek

Go-deepseek is a Go client for the [DeepSeek API](https://api-docs.deepseek.com/); supporting DeepSeek-V3, DeepSeek-R1, and more, with both streaming and non-streaming options. This **production-ready** client is actively maintained, with ongoing bug fixes and feature enhancements. 

![go-deepseek-design](https://github.com/user-attachments/assets/346806ad-7617-4690-b6b4-0b49707852d8)

## Demo

**30 seconds deepseek-demo ([code](https://github.com/go-deepseek/deepseek-demo/)):** 

left-side browser with **[chat.deepseek.com](https://chat.deepseek.com/)** **v/s** **[go-deepseek](https://github.com/go-deepseek/deepseek)** in right-side terminal.

https://github.com/user-attachments/assets/baa05145-a13c-460d-91ce-90129c5b32d7

## Install
```
go get github.com/go-deepseek/deepseek
```

## Usage

![go-deepseek-flow](https://github.com/user-attachments/assets/dfa6fc98-65f2-4a08-ab13-8c0732ac8302)

Here’s an example of sending a "Hello Deepseek!" message using `model=deepseek-chat` (**DeepSeek-V3 model**) and `stream=false`

```go
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
## Why yet another Go client?

We were looking for `Dedicated` & `Simple` Go Client for Deepseek API but we didn't find it so we built this one 😃

## What's special about this Go client?

- **Simple** – Check below, Go package structure with all exported entities. It is as simple as possible. Also, it's Go's idiomatic way - request is under request pkg, response is under response.

- **Complete** – Check below, It offers full support for all APIs, including their complete request and response payloads. (Note: Beta feature support coming soon.)

![go_pkg_structure](https://github.com/user-attachments/assets/729a2294-98fa-4f6e-b936-ae5eb1b624ff)

- **Reliable** – We have implemented numerous Go tests to ensure that all features work correctly at all times.

- **Performant** – Speed is crucial when working with AI models. We have optimized this client to deliver the fastest possible performance.

> [!IMPORTANT]  
> We know that sometimes **Deepseek API is down** but we won't let you down.
>
> We have **`Fake Client`** using which you can continue your development and testing even though Deepseek API is down.
>
> Use `fake.NewFakeCallbackClient(fake.Callbacks{})` / See example [`examples/81-fake-callback-client/main_test.go`](examples/81-fake-callback-client/main_test.go)

## Examples

Please check the [examples](examples/) directory, which showcases each feature of this client.

![examples](https://github.com/user-attachments/assets/032ff864-7da5-4b76-9484-836b52046614)

## Buy me a GitHub Star ⭐

If you like our work then please give github star to this repo. 😊
