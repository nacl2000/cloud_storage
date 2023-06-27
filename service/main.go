package main

import (
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/nacl2000/clould_storage/routers/login"
	"github.com/nacl2000/clould_storage/tools/path"
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
	htmlPathPattern := filepath.Join("frontend", "login", "*.html")
	router.LoadHTMLGlob(path.GetSourceCodePath(&htmlPathPattern))

	cookieStore := sessions.NewCookieStore([]byte("cloud_storage"))
	cookieStore.Options(sessions.Options{
		Path:    "/",
		MaxAge:  1800,
	})
	router.Use(sessions.Sessions("cloud_storage", cookieStore))

	v1 := router.Group("/v1")
	login.AddLoginRoutes(v1)
}
