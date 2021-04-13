package fake

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(mock MockService) http.Handler {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.GET("/*url", Handler(mock))
	engine.POST("/*url", Handler(mock))
	engine.PATCH("/*url", Handler(mock))
	engine.DELETE("/*url", Handler(mock))
	return engine
}
