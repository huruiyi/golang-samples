package main

import (
	"fmt"
	"os"
	"strconv"
)

func func0(name string) {
	fmt.Println("func0: I am", name)
}

func func1(name string, age int) {
	fmt.Println("func1: I am", name, "I am", age, "years old")
}

func func2(name1, name2 string) {
	fmt.Println("func2: first name:", name1, "second name:", name2)
}

func func3(num int) int {
	return num + 1
}

func func4(num int) (int, int) {
	return 123, 456
}

func func5(a int, b int) (c int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	c = a / b
	return c
}

func func6(a int, b int) (string, string) {
	return strconv.Itoa(a), strconv.Itoa(b)
}

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func mod(a int, b int) (c int, d int) {
	return a, b
}

func add_demo() {
	a := 1
	b := add1(a)
	fmt.Println(b, a)

	c := 1
	d := add2(&c)
	fmt.Println(d, c)

	fmt.Println(mod(1, 3))
}

/* 值传递*/
func SwapInt(num1, num2 int) {
	fmt.Println("SwapInt 前:", num1, num2)
	temp := num1
	num1 = num2
	num2 = temp
	fmt.Println("SwapInt 后:", num1, num2)
}

/* 引用传递*/
func SwapIntReference(num1 *int, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

func SwapDemo() {
	fmt.Println("**********************")
	var num1, num2 int = 123, 456
	fmt.Println(num1, num2)
	SwapInt(num1, num2)
	fmt.Println(num1, num2)
	fmt.Println("**********************")
	var a, b int = 123, 456
	fmt.Println("a,b = b,a 前", a, b)
	a, b = b, a
	fmt.Println("a,b = b,a 后", a, b)
	fmt.Println("**********************")
	var n1, n2 = 123, 456
	fmt.Println("SwapIntReference前", n1, n2)
	SwapIntReference(&n1, &n2)
	fmt.Println("SwapIntReference后", n1, n2)
}

func TestDefer() {
	file, err := os.Open("08_if.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}

func main() {
	println("*********************************")
	func0("小明")
	func1("小明", 123)
	func2("小明", "小红")
	println("*********************************")


	println("*********************************")
	var revalue0 int = func3(25)
	fmt.Println("func3: ",revalue0)
	revalue1 := func3(100)
	fmt.Println("func3: ",revalue1)
	println("*********************************")


	var a, b int = func4(234)
	fmt.Println("func4: ",a, b)

	num1 := 21
	num2 := 3
	var c = func5(num1, num2)
	fmt.Println("func5: ",c)

	s1, s2 := func6(num1, num2)
	println("func6: 输出: ", s1, s2)
}

/*
Ｃ语言中提供了地址运算符&来表示变量的地址。其一般形式为： & 变量名； 如&a变示变量a的地址，&b表示变量b的地址。
变量本身必须预先说明。设有指向整型变量的指针变量p，如要把整型变量a 的地址赋予p可以有以下两种方式：
(1)指针变量初始化的方法 int a;
int *p=&a;
(2)赋值语句的方法 int a;
int *p;
p=&a;
不允许把一个数赋予指针变量，故下面的赋值是错误的： int *p;p=1000; 被赋值的指针变量前不能再加“*”说明符，如写为*p=&a 也是错误的
-------------------------
• Go语言中string，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/
