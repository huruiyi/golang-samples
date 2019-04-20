// +build !appengine

package main

import (
	"net/http"

	_ "../hybridapplib"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
