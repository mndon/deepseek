package internal

import (
	"encoding/json"
)

type ErrorResponse struct {
	Error *Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   any    `json:"param"`
	Code    string `json:"code"`
}

func ParseError(errBody []byte) (*ErrorResponse, error) {
	errResp := &ErrorResponse{}
	err := json.Unmarshal(errBody, errResp)
	return errResp, err
}
