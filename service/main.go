package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/nacl_clould_storage/routers/login"
)

var router = gin.Default()

// Run will start the server
func main() {
	getRoutes()
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	login.AddLoginRoutes(v1)
}
