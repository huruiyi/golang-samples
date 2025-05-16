package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
)

type User struct {
	Name  string
	Email string
	ID    int
}

func userRouter(w http.ResponseWriter, r *http.Request) {
	ourUser := User{}
	ourUser.Name = "Bill Smith"
	ourUser.Email = "bill.smith@example.com"
	ourUser.ID = 100

	output, _ := yaml.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {
	//http://localhost:8090/user
	fmt.Println("Starting YAML server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8090", nil)

}
