package mock

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MockRoutes(mock MockService) http.Handler {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.GET("/*url", mockHandler(mock))
	engine.POST("/*url", mockHandler(mock))

	return engine
}
