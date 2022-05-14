package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", func(ctx *gin.Context) {
		genre := ctx.Query("genre")
		country := ctx.Query("country")

		// if condition {

		// }
		fmt.Println("DEB-", genre, country)
		if len(country) > 1 && len(genre) > 1 {
			ctx.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", genre, country)
		} else if len(country) > 1 {
			ctx.String(http.StatusOK, "Here result of query of movie with genre general in country %s", country)
		} else if len(genre) > 1 {
			ctx.String(http.StatusOK, "Here result of query of movie with genre %s", genre)
		} else {
			ctx.String(http.StatusOK, "Here result of query of movie with genre general")
		}
	})

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
