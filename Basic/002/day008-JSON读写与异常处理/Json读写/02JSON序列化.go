package main

import (
	"encoding/json"
	"fmt"
)

/*
√ 使用Person结构体描述谦哥并转JSON
√ 使用map[string]interface{}描述谦哥并转JSON
√ 使用map切片描述谦哥的小姨子们并转JSON
*/

type Person struct {
	Name  string
	Age   int
	Rmb   float64
	Sex   bool
	Hobby []string
}

/*使用Person结构体描述谦哥并转JSON*/
func main021() {
	person := Person{"于谦", 50, 123.45, true, []string{"抽烟", "喝酒", "烫头"}}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	jsonStr := string(jsonBytes)
	fmt.Println("序列化成功", jsonStr)

}

/*使用map[string]interface{}描述谦哥并转JSON*/
func main022() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 50
	dataMap["rmb"] = 123.45
	dataMap["sex"] = true
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	jsonBytes, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	jsonStr := string(jsonBytes)
	fmt.Println("序列化成功", jsonStr)
}

/*使用结构体切片描述谦哥的小姨子们并转JSON*/
func main() {
	p1 := Person{Name: "王钢蛋", Hobby: []string{"抽烟", "喝酒", "烫头"}}
	p2 := Person{Name: "王铁蛋", Hobby: []string{"抽中华", "喝牛栏山", "烫杀马特"}}
	p3 := Person{Name: "王铜蛋", Hobby: []string{"抽小熊猫", "喝剑南春", "烫鸡冠头"}}

	people := make([]Person, 0)
	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)

	jsonBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	jsonStr := string(jsonBytes)
	fmt.Println("序列化成功", jsonStr)
}
