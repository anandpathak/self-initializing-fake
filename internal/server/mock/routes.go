package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MockRoutes() http.Handler {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.GET("/*url", mockHandler() )

	return engine
}

func mockHandler()  gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Param("url")
		c.String(200, urlPath)
	}
}
