package main

import "fmt"

func main() {
	sum0()
	sum1()
}

func sum0() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum1() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
