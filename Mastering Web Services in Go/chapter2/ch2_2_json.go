package main

import (
	"./lib"
	"encoding/json"
	"fmt"
	"net/http"
)

func router2(w http.ResponseWriter, r *http.Request) {
	ourUser := lib.User{}
	ourUser.Name = "Bill Smith"
	ourUser.Email = "bill.smith@example.com"
	ourUser.ID = 100
	ourUser.First = "Ruiyi"
	ourUser.Last = "Hu"

	output, _ := json.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {

	fmt.Println("Starting JSON server")
	http.HandleFunc("/user", router2)
	http.ListenAndServe(":8090", nil)

}
