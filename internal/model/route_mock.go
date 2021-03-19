package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type RequestBodyForMock struct {
	ID       string            `json:"-"`
	Headers  map[string]string `json:"headers" binding:"required"`
	Request  interface{}       `json:"request" binding:"required"`
	Response interface{}       `json:"response" binding:"required"`
	URL      string            `json:"url" binding:"required"`
}

func (r RequestBodyForMock) GetHash() string {
	data := fmt.Sprintf("%v", r)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}