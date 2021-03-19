package service

import (
	"self_initializing_fake/internal/model"
	"self_initializing_fake/pkg/memorydb"
)

type Configure struct {
	DB memorydb.Store
}
type ConfigureService interface {
	Run(model.RequestBodyForMock) (string, error)
}

const (
	TableName = "mock_request"
)

func (c Configure) Run(request model.RequestBodyForMock) (string, error) {

	id := request.GetHash()
	request.ID = id

	if err := c.DB.Save(TableName, request); err != nil {
		return "", err
	}
	return id, nil
}
