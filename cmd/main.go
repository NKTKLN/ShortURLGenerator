package main

import (
	"github.com/NKTKLN/ShortURLGenerator/pkg/pages"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/NKTKLN/ShortURLGenerator/docs"
)

// @title Short URL Generator API
// @version 2.0
// @description This api will help you control your links from outside the site

// @contact.name API Support
// @contact.url https://nktkln.com
// @contact.email nktkln@nktkln.com

// @host localhost:8080
// @BasePath /api
func main() {
	r := gin.Default()
	// Connecting files to display pages
    r.LoadHTMLGlob("templates/*")
    r.Use(static.Serve("/static/", static.LocalFile("./static/", true)))
	// Pages
    r.NoRoute(pages.Error404)
	r.GET("/", pages.Home)
	r.POST("/", pages.Save)
	r.POST("/api/add", pages.ApiAdd)
	r.POST("/api/delete", pages.ApiDelete)
	r.GET("/api/archive", pages.ApiArchive)
	r.GET("/login", pages.SignIn)
	r.POST("/login", pages.Login)
	r.GET("/reset", pages.Reset)
	r.POST("/reset", pages.ResetPassword)
	r.GET("/reset/:hash", pages.ResetPasswordVerification)
	r.POST("/reset/:hash", pages.ResetPasswordSave)
	r.GET("/register", pages.SignUp)
	r.POST("/register", pages.Register)
	r.GET("/verification/:hash", pages.VerificationLogin)
	r.GET("/logout", pages.Logout)
	r.GET("/telegramLogin", pages.TelegramLogin)
	r.POST("/telegramLogin", pages.TelegramLoginDone)
	r.GET("/archive", pages.Archive)	
	r.POST("/archive", pages.Delete)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.GET("/:id", pages.Redirect)
	// Running the site on port 80
	r.Run(":80")
}
