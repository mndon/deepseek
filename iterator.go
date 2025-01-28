package deepseek

import (
	"bufio"
	"encoding/json"
	"io"
)

type MessageIterator struct {
	msgCh chan *DeepseekChatResponse
}

func NewMessageIterator(stream io.ReadCloser) *MessageIterator {
	iter := &MessageIterator{
		msgCh: make(chan *DeepseekChatResponse),
	}
	go iter.process(stream)
	return iter
}

func (m *MessageIterator) Next() *DeepseekChatResponse {
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
		bytes = TrimDataPrefix(bytes)
		if len(bytes) > 1 && bytes[0] == '[' {
			str := string(bytes)
			if str == "[DONE]" {
				close(m.msgCh)
				return
			}
		}
		chatResp := &DeepseekChatResponse{}
		err = json.Unmarshal(bytes, chatResp)
		if err != nil {
			panic(err)
		}
		m.msgCh <- chatResp
	}
}

func TrimDataPrefix(content []byte) []byte {
	trimIndex := 6
	if len(content) > trimIndex {
		return content[trimIndex:]
	}
	return content
}
