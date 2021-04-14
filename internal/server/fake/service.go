package fake

import (
	"errors"
	"fmt"
	"reflect"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/pkg/memorydb"
	"strings"
)

type Mock struct {
	DB memorydb.Store
}

type MockService interface {
	Run(model.TestDouble) (*model.TestDouble, error)
}

func (m Mock) Run(request model.TestDouble) (*model.TestDouble, error) {

	td, err := m.DB.FetchById("id", request.ID)
	if err != nil {
		return nil, err
	}

	fakeTestDouble, ok := td.(model.TestDouble)
	if !ok {
		return nil, errors.New(NoMockPresent)
	}
	if !isHeaderValid(request.Request.Header, fakeTestDouble.Request.Header) {
		return nil, errors.New(fmt.Sprintf(IncorrectHeader, request.Request.Header, fakeTestDouble.Request.Header))
	}

	if !fakeTestDouble.Request.IgnoreValidation {
		if !isRequestValid(request.Request.Body, fakeTestDouble.Request.Body) {
			fmt.Printf(IncorrectRequest, request.Request.Body, fakeTestDouble.Request.Body)
			return nil, errors.New(fmt.Sprintf(IncorrectRequest, request.Request.Body, fakeTestDouble.Request.Body))
		}
	}


	d := schedule{}
	response := <- d.Delay(fakeTestDouble)

	return &response, nil

}

func isRequestValid(request interface{}, mockedRequest interface{}) bool {
	return reflect.DeepEqual(request, mockedRequest)
}

func isHeaderValid(requestHeaders, cachedHeaders map[string][]string) bool {
	for headerName, headerValue := range cachedHeaders {

		result, found := requestHeaders[headerName]
		if !found {
			fmt.Printf("%v header is missing", headerName)
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
