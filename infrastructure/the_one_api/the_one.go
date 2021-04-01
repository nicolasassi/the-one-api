package the_one_api

import (
	"encoding/json"
	"time"
)

const (
	endpoint = "https://the-one-api.dev/v2/"
	Timeout  = 20 * time.Second
)

type response struct {
	Docs   []map[string]interface{} `json:"docs"`
	Total  int                      `json:"total"`
	Limit  int                      `json:"limit"`
	Offset int                      `json:"offset"`
	Page   int                      `json:"page"`
	Pages  int                      `json:"pages"`
}

func newResponse(data []byte) (*response, error) {
	var resp response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
