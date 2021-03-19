package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/internal/service"
)



func ConfigureHandler(configurationService service.ConfigureService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var request model.RequestBodyForMock
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Println("handler", request)
		if err := configurationService.Run(request);  err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"success": "ok"})

	}
}
