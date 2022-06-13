package http_utill

import (
	"encoding/json"
)

type response struct {
	Header responseHeader `json:"header"`
	Data   interface{}    `json:"data,omitempty"`
}

type responseHeader struct {
	IsSuccess bool     `json:"is_success"`
	Messages  []string `json:"messages"`
}

func ParseResponseToJSON(isSuccess bool, messages []string, data interface{}) ([]byte, error) {
	header := responseHeader{
		IsSuccess: isSuccess,
		Messages:  messages,
	}

	resp := response{
		Header: header,
		Data:   data,
	}

	bs, err := json.Marshal(&resp)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
