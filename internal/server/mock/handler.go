package mock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_initializing_fake/internal/model"
	"strings"
)

func mockHandler(mock MockService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var receivedRequest model.RequestBodyForMock
		var request interface{}

		receivedRequest.URL = fmt.Sprintf("%v", c.Request.URL)
		receivedRequest.Headers = c.Request.Header

		if err := c.ShouldBindJSON(&request); err == nil {
			receivedRequest.Request = request

		}
		receivedRequest.ID = receivedRequest.GetHash()

		r, err := mock.Run(receivedRequest)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for h, val := range r.Headers {
			v := strings.Join(val, ",")
			c.Header(h, v)
		}
		c.JSON(200, r.Response)
	}
}
