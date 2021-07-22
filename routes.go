// routes.go

package main

func initializaRoutes() {
	// Handle the Index route
	router.GET("/", showIndexPage)
}
