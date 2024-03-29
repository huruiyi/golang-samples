package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (ani *Animal) GetAge1() int {
	return ani.Age
}

func (ani Animal) GetAge2() int {
	return ani.Age
}

func main() {
	ani := new(Animal)
	ani.Age = 10
	//指针调用方法
	fmt.Println(ani.GetAge1())
	fmt.Println(ani.GetAge2())

	ani2 := Animal{"", 20}
	//非指针变量调用方法
	fmt.Println(ani2.GetAge1())
	fmt.Println(ani2.GetAge2())
}
