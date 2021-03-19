package service

import (
	"fmt"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/pkg/memorydb"
)

type Configure struct {
	DB memorydb.Store
}
type ConfigureService interface {
	Run(model.RequestBodyForMock) error
}

const (
	TableName="mock_request"

)
func (c Configure) Run( request model.RequestBodyForMock)  error{
	fmt.Printf("herheherhe")
	id := request.GetHash()
	request.ID = id

	if err := c.DB.Save(TableName, request);  err != nil {
			return err
	}
	fmt.Print("bobobo")
	fmt.Print(c.DB.FetchById("mock_request", "id", id ))
	return nil
}
