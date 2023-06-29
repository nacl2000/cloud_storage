package home

import (
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/nacl2000/clould_storage/tools/path"
)

func loadHomePage(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("username") == nil {
		c.Redirect(302, "/v1/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", "home")
}

func AddHomeRoutes(router *gin.RouterGroup) {
	homeRouter := router.Group("/home")
	homeRouter.GET("/", loadHomePage)

	homePageFrontedFileDirPath := filepath.Join("frontend","home")
	homeRouter.Static("/frontend", path.GetSourceCodePath(&homePageFrontedFileDirPath))
}
