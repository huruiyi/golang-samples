package main

import (
	"./lib"
	"fmt"
	"launchpad.net/goyaml"
	"net/http"
)

func userRouter(w http.ResponseWriter, r *http.Request) {
	ourUser := lib.User{}
	ourUser.Name = "Bill Smith"
	ourUser.Email = "bill.smith@example.com"
	ourUser.ID = 100

	output, _ := goyaml.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {

	fmt.Println("Starting YAML server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8080", nil)

}
