package route

import (
	. "github.com/gonic/middleware"
	. "github.com/gonic/handlers"
	"github.com/gin-gonic/gin"
)
var Router *gin.Engine

func InitializeRoutes() {
	Router.Use(SetUserStatus())
	Router.GET("/", ShowIndexPage)
	userRoutes := Router.Group("/u")
	{
		userRoutes.GET("/register", EnsureNotLoggedIn(), ShowRegistrationPage)
		userRoutes.POST("/register", EnsureNotLoggedIn(), Register)

		userRoutes.GET("/login", EnsureNotLoggedIn(), ShowLoginPage)
		userRoutes.POST("/login", EnsureNotLoggedIn(), PerformLogin)

		userRoutes.GET("/logout", EnsureLoggedIn(), Logout)

	}

	articleRoutes := Router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", GetArticle)
		articleRoutes.GET("/create", EnsureLoggedIn(), ShowArticleCreatePage)
		articleRoutes.POST("/create", EnsureLoggedIn(), CreateArticle)
	}

}