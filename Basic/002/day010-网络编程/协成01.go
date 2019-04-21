package main

import (
	"fmt"
	"time"
)

func numbers(str string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}

func test01() {
	numbers("a")
	numbers("b")
	numbers("c")
}

func test02() {
	go numbers("a")
	go numbers("b")
	go numbers("c")

}
func main() {
	test02()
}
