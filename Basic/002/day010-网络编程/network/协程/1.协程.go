package main

import (
	"fmt"
	"time"
)

func numbers(str string) {
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}

func main() {
	go numbers("b")
	go numbers("c")
	numbers("a")
}
