package fake

import (
	"fmt"
	"net/http"
	"self_initializing_fake/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func Handler(mock MockService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var receivedRequest model.TestDouble

		receivedRequest.URL = fmt.Sprintf("%v", c.Request.URL)
		receivedRequest.Request.Header = c.Request.Header

		var rr interface{}
		if err := c.ShouldBindJSON(&rr); err == nil {
			receivedRequest.Request.Body = rr
		}

		receivedRequest.ID = receivedRequest.GetHash()

		fakeResponse, err := mock.Run(receivedRequest)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		setHeaders(c, fakeResponse.Response.Header)
		c.JSON(200, fakeResponse.Response)
	}
}

func setHeaders(c *gin.Context, headers map[string][]string) {
	for h, val := range headers {
		v := strings.Join(val, ",")
		c.Header(h, v)
	}
}
