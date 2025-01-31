package deepseek

import (
	"bufio"
	"encoding/json"
	"io"
)

type MessageIterator struct {
	msgCh chan *ChatCompletionsResponse
}

func NewMessageIterator(stream io.ReadCloser) *MessageIterator {
	iter := &MessageIterator{
		msgCh: make(chan *ChatCompletionsResponse),
	}
	go iter.process(stream)
	return iter
}

func (m *MessageIterator) Next() *ChatCompletionsResponse {
	return <-m.msgCh
}

func (m *MessageIterator) process(stream io.ReadCloser) {
	defer stream.Close()
	reader := bufio.NewReader(stream)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		if len(bytes) <= 1 {
			continue
		}
		bytes = trimDataPrefix(bytes)
		if len(bytes) > 1 && bytes[0] == '[' {
			str := string(bytes)
			if str == "[DONE]" {
				close(m.msgCh)
				return
			}
		}
		chatResp := &ChatCompletionsResponse{}
		err = json.Unmarshal(bytes, chatResp)
		if err != nil {
			panic(err)
		}
		m.msgCh <- chatResp
	}
}

func trimDataPrefix(content []byte) []byte {
	trimIndex := 6
	if len(content) > trimIndex {
		return content[trimIndex:]
	}
	return content
}
