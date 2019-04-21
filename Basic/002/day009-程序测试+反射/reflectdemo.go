package main

import (
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Hobby  [] string
	IsMale bool
}

func main() {

	person := Person{Name: "羡慕", Age: 12, IsMale: false}
	name := reflect.TypeOf(person).Name()
	println("1::", name)

	nums := reflect.TypeOf(person).NumField()
	println("参数个数：", nums)
	for i := 0; i < nums; i++ {
		field := reflect.TypeOf(person).Field(i)
		println(field.Name)
	}
}
