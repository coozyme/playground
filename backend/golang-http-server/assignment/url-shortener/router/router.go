package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	// TODO: replace this
	r := gin.Default()
	r.GET("/:link", urlHandler.Get)

	return r
}
