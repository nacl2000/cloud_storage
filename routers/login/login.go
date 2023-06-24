package login

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddLoginRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "login")
	})
}
