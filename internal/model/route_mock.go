package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Header map[string]interface{}

type RequestBodyForMock struct {
	ID       string      `json:"-"`
	Headers  Header      `json:"headers" binding:"required"`
	Request  interface{} `json:"request" binding:"required"`
	Response interface{} `json:"response" binding:"required"`
	URL      string      `json:"url" binding:"required"`
}

func (r RequestBodyForMock) GetHash() string {
	data := fmt.Sprintf("%s", r.URL)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
