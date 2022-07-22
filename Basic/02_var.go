package main

import (
	"errors"
	"fmt"
	"reflect"
)

var x, y int = 123, 456
var c, d int = 1, 2
var e, f = 123, "hello"

//全局变量,不使用不会报错
var cc string = "dd"
var num int = 1244

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

var (
	// 这种因式分解关键字的写法一般用于声明全局变量
	a        int    = 100
	b        bool   = false
	nickname string = "dd"
)

const (
	//常量定义，不可修改
	version = "1.0"
)

// int、float、bool 和 string 都属于值类型，变量直接指向存在内存中的值：
func ValueTypeDemo() {}

func printVar() {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("cc:", cc)
	fmt.Println("num:", num)
	fmt.Println("version:", version)
	fmt.Println("nickname:", nickname)
}

func DeclareDemo() {
	var i int
	i = 1
	var j int = 2
	var k = 3
	l := 4
	var m int = 5
	var n string = "6"
	var (
		o int    = 7
		p string = "8"
	)
	println(i, j, k, l)
	println(m, n)
	println(o, p)
}

func ErrorDemo() {
	err := errors.New("测试肺胀")
	if err != nil {
		fmt.Print("error info " + err.Error())
	}
}

func PointDemo() {
	var a int = 123
	var p *int = &a
	*p = 456
	println(a)
}

func Demo0() {
	fmt.Println("***********************")
	e, f := 456, "world"
	g, h := 123, "hello"
	println(x, y)
	println(a, b, c, d, e, f)
	println(g, h)

	fmt.Println("***********************")
	var (
		xa int
		xb bool
	)
	xa = 123
	xb = false
	fmt.Println(xa, xb)
	fmt.Println("***********************")
}

func Demo1() {
	// 区别? 1.此种方式不用会报错：c declared and not used 2.仅在函数内使用
	cc := "ddd"
	fmt.Print(cc)

	cc = "aaa"
	fmt.Print(cc)
}

func Demo2() {
	var i = `ddd`
	//只能用到函数内部
	j := 0
	fmt.Printf("%sImy%d\n", i, j)
	nickname = "dd0"
	fmt.Print(nickname)
	var a, b, c = 1, 2, 3
	fmt.Print(a, b, c)
	var d, _ = 1.22, 33 //特殊变量
	fmt.Print(d)
}

func Demo3() {
	i1 := 123
	fmt.Println(reflect.TypeOf(i1))

	i2 := 123.123
	fmt.Println(reflect.TypeOf(i2))

	i3 := -1
	fmt.Println(reflect.TypeOf(i3))

	s1 := "123"
	fmt.Println(reflect.TypeOf(s1))

	var num int = 123
	fmt.Println(num)
}

func main() {
	printVar()
}
