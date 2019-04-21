package main

import "fmt"

type Person struct {
	//封装结构体的属性
	name  string
	age   int
	sex   bool
	hobby []string
}

/*
封装结构体的方法
- 无论方法的主语定义为值类型还是指针类型，对象值和对象指针都能够正常访问
- 通常会将主语定义为指针类型——毕竟西门鹤的副本吃了饭，肉不会长到西门鹤本人身上去
*/
func (p *Person) Eat() () {
	fmt.Printf("%s爱饕餮\n", p.name)
}
func (p *Person) Drink() {
	fmt.Printf("%s爱喝酒\n", p.name)
}
func (p *Person) Love() {
	fmt.Printf("%s是有感情的\n", p.name)
	p.age -= 1
}
func (p *Person) SelfIntroduce() {
	fmt.Printf("我是%s，今年%d岁了\n", p.name, p.age)
}

//值传递传递的是对象的副本
func MakeHimLove(p Person) {
	fmt.Printf("p的地址是%p\n",&p)
	p.Love()
}

//引用传递传递的是对象的真身
func MakeHisPtrLove(p *Person) {
	fmt.Printf("p的地址是%p\n",p)
	p.Love()
}

func main011() {
	////创建空白的Person对象（object）/实例(instance)
	//rangge := Person{}
	//fmt.Printf("%T\n",rangge)

	////根据实例访问属性和方法
	//rangge.name = "西门阿让"
	//rangge.Eat()
	//rangge.Drink()
	//rangge.Love()

	//通过指针访问属性和方法
	rangge := &Person{}
	fmt.Printf("%T\n",rangge)
	rangge.name = "西门阿让"
	rangge.Eat()
	rangge.Drink()
	rangge.Love()
}

func main012() {
	//创建对象时给指定属性赋值
	//rangge := Person{name:"西门阿让",sex:true,age:8}

	//创建对象定义属性的顺序给所有属性赋值
	//rangge := Person{"西门阿让", 8, true, []string{"撸代码", "完美的撸代码"}}
	//rangge := Person{"西门阿让", 8, true, []string{"撸代码", "完美的撸代码"}}

	rangge := new(Person)
	rangge.name = "西门爱让"
	rangge.age = 8

	rangge.Eat()
	rangge.Drink()
	rangge.Love()
	rangge.SelfIntroduce()
}

func main() {
	rangge := Person{"西门阿让", 8, true, []string{"撸代码", "完美的撸代码"}}
	fmt.Printf("让哥的真身地址是%p\n",&rangge)//

	//要求传递值就必须传递值
	//MakeHimLove(rangge)

	//要求传递指针就必须传递指针（指针/地址/引用传递）
	MakeHisPtrLove(&rangge)

	////值传递传递的是副本，引用传递传递的才是真身
	//for i := 0; i < 7; i++ {
	//	//MakeHimLove(rangge)
	//	MakeHisPtrLove(&rangge)
	//}
	fmt.Printf("暴风雨过后让哥的年龄是%d\n",rangge.age)
}
