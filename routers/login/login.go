package login

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/clould_storage/dal"
	"github.com/nacl2000/clould_storage/tools/calculte"
	"net/http"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Initial(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "login")
}

func Login(c *gin.Context) {
	var requester user
	if err := c.BindJSON(&requester); err != nil {
		c.String(404, "Login failed!")
		return
	}
	if requester.Username == "" || requester.Password == "" {
		c.String(404, "Username or password should not be empty.")
		return
	}
	if !authUser(requester) {
		c.String(http.StatusUnauthorized, "Username or passwd is error")
		return
	}

	session := sessions.Default(c)
	session.Set("username", requester.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"username": session.Get("username")})
}

func authUser(userInfo user) bool {
	passwd := dal.QueryPasswdByMd5(userInfo.Username)
	inputPasswd := calculte.GetSaltingPwd(userInfo.Password)
	return strings.ToUpper(passwd) == strings.ToUpper(inputPasswd)
}
