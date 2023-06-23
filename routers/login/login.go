package login

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddLoginRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/login")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login")
	})
}
