package mock

import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/internal/service"

	"github.com/gin-gonic/gin"
)

func MockRoutes(mock service.MockService) http.Handler {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.GET("/*url", mockHandler(mock))
	engine.POST("/*url", mockHandler(mock))

	return engine
}

func mockHandler(mock service.MockService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var receivedRequest model.RequestBodyForMock
		var headers model.Header
		var request interface{}

		receivedRequest.URL = fmt.Sprintf("%v", c.Request.URL)

		h, err := json.Marshal(c.Request.Header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		_ = json.Unmarshal(h, &headers)

		receivedRequest.Headers = headers

		if err := c.ShouldBindJSON(&request); err != nil {
			fmt.Printf("request parsing invalid : %v", request)
		} else {
			fmt.Printf("request here %v", request)
			receivedRequest.Request = request
		}

		receivedRequest.ID = receivedRequest.GetHash()
		var r *model.RequestBodyForMock
		if r, err = mock.Run(receivedRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for h, val := range r.Headers {
			v := fmt.Sprintf("%v", val)
			c.Header(h, v)
		}
		c.JSON(200, r.Response)
	}
}
