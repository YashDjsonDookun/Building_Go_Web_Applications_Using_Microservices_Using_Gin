// routes.go
package main

func initializaRoutes() {
	// Handle the Index route
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)
}
