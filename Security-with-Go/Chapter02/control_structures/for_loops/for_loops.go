package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	n := 5
	for n < 10 {
		fmt.Println(n)
		n++
	}
}
