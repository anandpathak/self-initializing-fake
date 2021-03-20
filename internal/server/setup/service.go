package setup

import (
	"self_initializing_fake/internal/model"
	"self_initializing_fake/pkg/memorydb"
)

type Service struct {
	DB memorydb.Store
}
type Runner interface {
	Run(model.TestDouble) (string, error)
}

const (
	TableName = "mock_request"
)

func (c Service) Run(request model.TestDouble) (string, error) {

	id := request.GetHash()
	request.ID = id

	if err := c.DB.Save(request); err != nil {
		return "", err
	}
	return id, nil
}
