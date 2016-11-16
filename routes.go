package main

func initializeRoutes() {
	router.GET("/", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/register", register)
	}


}