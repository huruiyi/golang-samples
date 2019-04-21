package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	resp, err := http.Post("https://httpbin.org/post?name=jack&teacher=john",
		"application/x-www-form-urlencoded",
		strings.NewReader("姓名=接客&师承=约汉&123=abc"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes)) //打印字节集

}
