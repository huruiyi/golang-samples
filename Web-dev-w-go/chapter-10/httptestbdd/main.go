package main

import (
	"net/http"

	"./lib"
)

func main() {
	routers := lib.SetUserRoutes()
	http.ListenAndServe(":8080", routers)
}
