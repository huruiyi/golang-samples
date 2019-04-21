package main

import (
	"math"
	"errors"
	"fmt"
)

/*
√ 求玩具球的体积
√ 半径如果不在取值范围内，温和地返回【结果-错误对】
*/
func GetBallVolumn2(radius float64) (volumn float64,err error) {
	if radius < 0 {
		//报出恐慌，程序死亡，以严厉的方式提示调用者，半径不能为负数
		panic("半径不能为负数")
	}

	//超出了工厂的生产能力，返回错误
	if radius < 5 || radius > 50{
		err = errors.New("合法半径为[5,50]")
		return
	}

	return (4.0 / 3) * math.Pi * math.Pow(radius, 3),nil
}

func main() {
	volumn, err := GetBallVolumn2(60)
	if err != nil{
		fmt.Println("获取体积失败，err=",err)
	}else{
		fmt.Println("小球的体积是",volumn)
	}
}
