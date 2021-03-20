package setup

import (
	"net/http"
	"self_initializing_fake/internal/model"

	"github.com/gin-gonic/gin"
)

func Handler(configurationService Runner) gin.HandlerFunc {

	return func(c *gin.Context) {
		var request model.TestDouble

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		}

		result, err := configurationService.Run(request)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "data": result})

	}
}
