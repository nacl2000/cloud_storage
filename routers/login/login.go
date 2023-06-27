package login

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

type user struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
}

func login(c *gin.Context) {
	var requester user
	if err := c.BindJSON(&requester); err != nil {
		c.String(404, "Login failed!")
		return
	}
	if requester.Username ==  "" ||  requester.Password == "" {
		c.String(404, "Username or password should not be empty.")
		return
	}
	session := sessions.Default(c)
	session.Set("username", requester.Username)
	session.Save()
	c.JSON(200, gin.H{"username": session.Get("username")})
}

func AddLoginRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "login")
	})
	loginRouter.POST("/api", login)
}
