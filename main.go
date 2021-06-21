package main

import (
	"go_web_store/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
