package main

func initializeRoutes() {

	// Используйте промежуточное программное обеспечение setUserStatus для каждого маршрута, чтобы установить флаг
	// указание, был ли запрос от аутентифицированного пользователя или нет
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/NickName")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:NickName_id", getArticle)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
