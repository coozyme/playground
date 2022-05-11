package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	// TODO: replace this
	router := gin.Default()

	router.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	router.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	router.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	router.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	router.PATCH("/patch", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PATCH",
		})
	})
	router.HEAD("/head", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "HEAD",
		})
	})
	router.OPTIONS("/options", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OPTIONS",
		})
	})
	return router
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
