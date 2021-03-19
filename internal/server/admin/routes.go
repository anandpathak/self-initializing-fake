package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self_initializing_fake/internal/service"
)


func AdminRoutes(configurationService service.ConfigureService) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	adminRouteGroup := engine.Group("/admin")
	{
		adminRouteGroup.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		adminRouteGroup.POST("/configure",  ConfigureHandler(configurationService))
	}
	return engine
}


