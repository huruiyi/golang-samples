package main

import (
	"./lib"
	"encoding/xml"
	"fmt"
	"net/http"
)

func router3(w http.ResponseWriter, r *http.Request) {
	ourUser := lib.User{}
	ourUser.Name = "Bill Smith"
	ourUser.Email = "bill.smith@example.com"
	ourUser.ID = 100

	output, _ := xml.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {

	fmt.Println("Starting JSON server")
	http.HandleFunc("/user", router3)
	http.ListenAndServe(":8090", nil)

}
