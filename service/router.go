package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/clould_storage/routers/login"
	"github.com/nacl2000/clould_storage/tools/path"
	"path/filepath"
)

func register(router *gin.Engine) {
	loadHtml(router)

	routerV1 := router.Group("/v1")

	loginRoutes(routerV1)

	userRoutes(routerV1)
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func loadHtml(router *gin.Engine) {
	htmlPathPattern := filepath.Join("frontend", "login", "*.html")
	router.LoadHTMLGlob(path.GetSourceCodePath(&htmlPathPattern))

	cookieStore := sessions.NewCookieStore([]byte("cloud_storage"))
	cookieStore.Options(sessions.Options{
		Path:   "/",
		MaxAge: 1800,
	})
	router.Use(sessions.Sessions("cloud_storage", cookieStore))
}

// login routes create route about all login things
func loginRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")

	loginRouter.GET("/", login.Initial)
	loginRouter.POST("/api", login.Login)
}

func userRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userRouter.POST("/register", login.UserRegister)
}
