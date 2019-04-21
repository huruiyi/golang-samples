package main

import (
	"math"
	"fmt"
)

/*
求玩具球的体积
半径如果为负数，直接panic
处理恐慌
半径如果不在取值范围内，温和地返回【结果-错误对】
√ 无论是panic还是返回，都使用自定义错误InvalidRadiusError的实例来操作
*/
func GetBallVolumn3(radius float64) (volumn float64,err error) {
	if radius < 0 {
		//报出恐慌，程序死亡，以严厉的方式提示调用者，半径不能为负数
		//panic(&InvalidRadiusError{radius,5,50})
		panic(NewInvalidRadiusError(radius))
	}

	//超出了工厂的生产能力，返回错误
	if radius < 5 || radius > 50{
		//err = errors.New("合法半径为[5,50]")
		//err = &InvalidRadiusError{}
		//err = &InvalidRadiusError{radius,5,50}
		err = NewInvalidRadiusError(radius)
		return
	}

	return (4.0 / 3) * math.Pi * math.Pow(radius, 3),nil
}

/*实现系统error接口，来自定义异常*/
type InvalidRadiusError struct {
	Value     float64
	MinRadius float64
	MaxRadius float64
}
func (ire *InvalidRadiusError)Error() string{
	str := fmt.Sprintf("%.2f是非法半径，合法半径的范围是[%.2f,%.2f]\n", ire.Value, ire.MinRadius, ire.MaxRadius)
	return str
}

/*创建异常实例的工厂方法*/
func NewInvalidRadiusError(value float64) *InvalidRadiusError {
	//封装起重复调用的逻辑
	ire := new(InvalidRadiusError)
	ire.MinRadius = 5
	ire.MaxRadius = 50

	//接收动态变化的核心数据
	ire.Value = value
	return ire
}

func main() {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("致死的原因是：",err)
		}
	}()

	volumn, err := GetBallVolumn3(-60)
	if err != nil{
		fmt.Println("获取体积失败，err=",err)
	}else{
		fmt.Println("小球的体积是",volumn)
	}
}
