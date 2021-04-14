package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Header map[string][]string

type TestDouble struct {
	ID       string   `json:"-"`
	Request  Request  `json:"request" binding:"required"`
	Response Response `json:"response" binding:"required"`
	URL      string   `json:"url" binding:"required"`
}

type Request struct {
	Header map[string][]string `json:"headers"`
	Body   interface{}         `json:"body"`
	IgnoreValidation bool `json:"ignore_validation"`
}

type Response struct {
	Header     map[string][]string `json:"headers"`
	Body       interface{}         `json:"body" binding:"required"`
	StatusCode int                 `json:"status_code" binding:"required"`
	Latency    int                 `json:"latency"`
}

func (r TestDouble) GetHash() string {
	data := fmt.Sprintf("%s", r.URL)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
