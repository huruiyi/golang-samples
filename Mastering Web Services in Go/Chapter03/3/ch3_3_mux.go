package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	println(router)
}
