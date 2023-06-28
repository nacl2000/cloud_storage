package main

import (
	"github.com/gin-gonic/gin"
)

// Run will start the server
func main() {
	var router = gin.Default()
	register(router)
	
	router.Run(":8080")
}
