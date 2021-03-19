package admin

import (
	"fmt"
	"net/http"
	"self_initializing_fake/internal/model"
	"self_initializing_fake/internal/service"

	"github.com/gin-gonic/gin"
)

func ConfigureHandler(configurationService service.ConfigureService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var request model.RequestBodyForMock
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Println("handler", request)
		result, err := configurationService.Run(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"success": result})

	}
}
