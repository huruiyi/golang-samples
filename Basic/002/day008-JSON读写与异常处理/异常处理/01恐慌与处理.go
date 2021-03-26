package main

import (
	"fmt"
	"math"
)

/*演示几个恐慌：对空指针取值、下标越界、向空map中添加键值对*/
func main011() {

	//对空指针取值
	//var intPtr *int
	//fmt.Println(*intPtr)

	//下标越界
	//mSlice := make([]int, 0)
	//mSlice = append(mSlice, 1, 2, 3, 4, 5)
	//fmt.Println(mSlice[10])

	//向空map中添加键值对
	var mMap map[string]interface{}
	mMap["name"] = "谦哥"

}

/*
√ 求玩具球的体积
√ 半径如果为负数，直接panic
√ 处理恐慌
*/
func GetBallVolumn(radius float64) float64 {
	if radius < 0 {
		//报出恐慌，程序死亡，以严厉的方式提示调用者，半径不能为负数
		panic("半径不能为负数")
	}
	return (4.0 / 3) * math.Pi * math.Pow(radius, 3)
}

func main() {
	//恐慌造成函数的异常返回，在返回之前的defer任务中，对恐慌进行处理
	defer func() {
		//处理恐慌
		//从恐慌中复生，捕获造成恐慌的元凶，进行处理
		if panicReason := recover(); panicReason != nil {
			fmt.Println(panicReason)
			fmt.Println("送钱送房子送美女2")
		}
	}()

	//程序因恐慌，卒！
	volumn := GetBallVolumn(-1)

	fmt.Println("送钱送房子送美女")
	fmt.Println("小球的体积是", volumn)
}
