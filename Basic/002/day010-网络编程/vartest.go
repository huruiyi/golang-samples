package main

import (
	"reflect"
)

var a int

func changeArr(arr []int) {
	arr[0] = 111
	println("changeArr:", arr)
}

func main() {

	var b int
	b = 12
	println(b)

	var a = 123
	println(a)

	println(&a)

	var i = &a
	println(i)

	println(reflect.TypeOf(i).Align())

	i2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(len(i2))

	i3 := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(i3); i++ {
		print(i3[i], ",")
	}
	println()
	changeArr(i3)

	for i := 0; i < len(i3); i++ {
		print(i3[i], ",")
	}
	println()

	var str = "Hello World,世界你好"
	println(str)

	for i := 0; i < len(str); i++ {
		println(string(str[i]))
	}
}
