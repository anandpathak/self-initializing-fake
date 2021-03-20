package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(configurationService Runner) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	adminRouteGroup := engine.Group("/setup")
	{
		adminRouteGroup.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		adminRouteGroup.POST("/fake_route", Handler(configurationService))
	}
	return engine
}
