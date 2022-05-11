package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	// TODO: replace this
	// id := c.

	return func(c *gin.Context) {}
}

func GetRouter() *gin.Engine {
	// TODO: replace this
	r := gin.Default()

	r.GET("/profile/:id", ProfileHandler)
	return &gin.Engine{}
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
