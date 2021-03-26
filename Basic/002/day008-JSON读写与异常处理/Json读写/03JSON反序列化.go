package main

import (
	"encoding/json"
	"fmt"
)

/*
√ 将谦哥老婆的JSON信息转换为map
√ 将谦哥老婆的JSON信息转换为结构体
√ 将谦哥小姨子们的JSON信息转换为map切片
*/

/*将谦哥的JSON信息转换为map*/
func main031() {
	jsonStr := `{"Name":"王钢蛋","Age":0,"Rmb":0,"Sex":false,"Hobby":["抽烟","喝酒","烫头"]}`
	jsonBytes := []byte(jsonStr)
	dataMap := make(map[string]interface{})

	err := json.Unmarshal(jsonBytes, &dataMap)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	fmt.Println("反序列化成功：", dataMap)
}

/*将谦哥的JSON信息转换为结构体*/
func main032() {

	type Person struct {
		Name  string
		Age   int
		Rmb   float64
		Sex   bool
		Hobby []string
	}

	jsonStr := `{"Name":"王钢蛋","Age":0,"Rmb":0,"Sex":false,"Hobby":["抽烟","喝酒","烫头"]}`
	jsonBytes := []byte(jsonStr)
	person := Person{}

	err := json.Unmarshal(jsonBytes, &person)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	fmt.Printf("反序列化成功：%+v\n", person)
}

/*将谦哥小姨子们的JSON信息转换为map切片*/
func main033() {
	jsonStr := `[{"Name":"王钢蛋","Age":0,"Rmb":0,"Sex":false,"Hobby":["抽烟","喝酒","烫头"]},{"Name":"王铁蛋","Age":0,"Rmb":0,"Sex":false,"Hobby":["抽中华","喝牛栏山","烫杀马特"]},{"Name":"王铜蛋","Age":0,"Rmb":0,"Sex":false,"Hobby":["抽小熊猫","喝剑南春","烫鸡冠头"]}]`
	jsonBytes := []byte(jsonStr)
	dataMap := make([]map[string]interface{}, 0)

	err := json.Unmarshal(jsonBytes, &dataMap)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
		return
	}
	fmt.Printf("反序列化成功：%+v\n", dataMap)
}
