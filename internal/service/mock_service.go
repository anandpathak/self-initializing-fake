package service

import (
	"errors"
	"fmt"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/pkg/memorydb"
	"self_initializing_fake/pkg/utils"
)

type Mock struct {
	DB memorydb.Store
}

type MockService interface {
	Run(model.RequestBodyForMock) (*model.RequestBodyForMock, error)
}


func (m Mock) Run(request model.RequestBodyForMock) (*model.RequestBodyForMock, error) {

	getResponse, err := m.DB.FetchById(TableName, "id",  request.ID)
	if err != nil {
		return nil, err
	}

	r, ok := getResponse.(model.RequestBodyForMock); if !ok {
		return nil,errors.New("mock not present")
	}
	if !isMockValid(request, r) {
		return nil, errors.New("request/header not matching")
	}

	return &r, nil
}

func isMockValid(request, cachedMock model.RequestBodyForMock ) bool {

	fmt.Printf("\n\n%v == %v\n\n", request.Headers, cachedMock.Headers)
	if !isHeaderValid(request.Headers, cachedMock.Headers) {
		return false
	}
	r := fmt.Sprintf("%v", request.Request)
	cr := fmt.Sprintf("%v", cachedMock.Request)

	if r != cr {
		return false
	}
	return true
}

func isHeaderValid(requestHeaders, cachedHeaders model.Header) bool {
	for h, _ := range requestHeaders {
		if !ignorableHeaders(h) && !utils.IncludesKey(cachedHeaders, h)  {
			fmt.Printf("%v is missing", h)
			return false
		}
	}
	return true
}
func ignorableHeaders(header string)  bool{
	ih :=  map[string]bool{
		"Accept": false,
		"Content-Type": false,
		"Content-Length": false,
	}
	if _, found := ih[header]; found {
		return true
	}
	return false
}