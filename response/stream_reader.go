package response

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
)

const KEEP_ALIVE = `: keep-alive`

const KEEP_ALIVE_LEN = len(KEEP_ALIVE)

type StreamReader interface {
	Read() (*ChatCompletionsResponse, error)
}

type streamReader struct {
	respCh chan *streamResponse
}

type streamResponse struct {
	chatResp *ChatCompletionsResponse
	error
}

func NewStreamReader(stream io.ReadCloser) StreamReader {
	iter := &streamReader{
		respCh: make(chan *streamResponse),
	}
	go iter.process(stream)
	return iter
}

func (m *streamReader) Read() (*ChatCompletionsResponse, error) {
	resp := <-m.respCh
	return resp.chatResp, resp.error
}

func (m *streamReader) process(stream io.ReadCloser) {
	defer stream.Close()
	reader := bufio.NewReader(stream)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			m.respCh <- &streamResponse{nil, err}
			close(m.respCh)
			return
		}
		if len(bytes) <= 1 {
			continue
		}
		chatResp, err := processResponse(bytes)
		if err != nil {
			m.respCh <- &streamResponse{nil, err}
			close(m.respCh)
			return
		}
		m.respCh <- &streamResponse{chatResp, err}
	}
}

func processResponse(bytes []byte) (*ChatCompletionsResponse, error) {
	// handle keep-alive response
	if len(bytes) == KEEP_ALIVE_LEN {
		if string(bytes) == KEEP_ALIVE {
			err := errors.New("err: service unavailable")
			return nil, err
		}
	}

	// handle response end
	bytes = trimDataPrefix(bytes)
	if len(bytes) > 1 && bytes[0] == '[' {
		str := string(bytes)
		if str == "[DONE]" {
			return nil, io.EOF // io.EOF to indicate end
		}
	}

	// parse response
	chatResp := &ChatCompletionsResponse{}
	err := json.Unmarshal(bytes, chatResp)
	return chatResp, err
}

func trimDataPrefix(content []byte) []byte {
	trimIndex := 6
	if len(content) > trimIndex {
		return content[trimIndex:]
	}
	return content
}
