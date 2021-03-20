package fake

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(mock MockService) http.Handler {

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.GET("/*url", Handler(mock))
	engine.POST("/*url", Handler(mock))

	return engine
}
