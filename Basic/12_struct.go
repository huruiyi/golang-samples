package main

import "fmt"

type human struct {
	name string
	age  int
}

type teacher struct {
	human
	sex int
}

func (t teacher) print() {
	fmt.Print("name:", t.human.name, "\t")
	fmt.Print("age:", t.human.age, "\t")
	fmt.Print("sex:", t.sex, "\t")
	fmt.Println()
}

func (t *teacher) printLn() {
	fmt.Println("name:", t.human.name)
	fmt.Println("age:", t.human.age)
	fmt.Println("sex:", t.sex)
}

func main() {
	t := teacher{human: human{name: "joe", age: 18}, sex: 1}
	t.print()
	t.printLn()
}
