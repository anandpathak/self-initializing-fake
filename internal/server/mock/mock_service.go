package mock

import (
	"errors"
	"fmt"
	"reflect"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/internal/server/admin"
	"self_initializing_fake/pkg/memorydb"
	"strings"
)

type Mock struct {
	DB memorydb.Store
}

type MockService interface {
	Run(model.RequestBodyForMock) (*model.RequestBodyForMock, error)
}

func (m Mock) Run(request model.RequestBodyForMock) (*model.RequestBodyForMock, error) {

	getResponse, err := m.DB.FetchById(admin.TableName, "id", request.ID)
	if err != nil {
		return nil, err
	}

	r, ok := getResponse.(model.RequestBodyForMock)
	if !ok {
		return nil, errors.New(NoMockPresent)
	}
	if !isHeaderValid(request.Headers, r.Headers) {
		return nil, errors.New(fmt.Sprintf(IncorrectHeader, request.Headers, r.Headers))
	}

	if !isRequestValid(request.Request, r.Request) {
		return nil, errors.New(fmt.Sprintf(IncorrectRequest, request.Request, r.Request))
	}

	return &r, nil
}

func isRequestValid(request interface{}, mockedRequest interface{}) bool {
	return reflect.DeepEqual(request, mockedRequest)
}

func isHeaderValid(requestHeaders, cachedHeaders map[string][]string) bool {
	for headerName, headerValue := range cachedHeaders {

		result, found := requestHeaders[headerName]
		if !found {
			fmt.Printf("%v is missing", headerName)
			return false
		}
		r := strings.Join(result, ",")
		h := strings.Join(headerValue, ",")
		if r != h {
			fmt.Printf("failing keys, %s, %s, %v %v", result[0], headerValue[0], result, headerValue[0])
			return false
		}
	}
	return true
}
