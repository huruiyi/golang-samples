package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
√ 使用JSON包的编码器编码谦嫂的信息为JSON文件
√ 编码谦哥八大姨结构体切片为JSON文件
√ 解码《谦嫂.json》为map
√ 解码《谦嫂.json》为结构体
√ 解码谦哥八大姨JSON文件为结构体切片
*/

/*使用JSON包的编码器编码谦嫂的信息为JSON文件*/
func main041() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "于谦"
	dataMap["age"] = 50
	dataMap["rmb"] = 123.45
	dataMap["sex"] = true
	dataMap["hobby"] = []string{"抽烟", "喝酒", "烫头"}

	jsonFile, _ := os.OpenFile("C:/Tencent1803/W2/day008_JSON读写与异常处理/file/于谦.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err := encoder.Encode(dataMap)
	if err != nil {
		fmt.Println("编码为JSON文件失败,err=", err)
		return
	}
	fmt.Println("编码成功！")
}

/*编码谦哥八大姨结构体切片为JSON文件*/
func main042() {

	type Person struct {
		Name  string
		Age   int
		Rmb   float64
		Sex   bool
		Hobby []string
	}
	p1 := Person{Name: "王钢蛋", Hobby: []string{"抽烟", "喝酒", "烫头"}}
	p2 := Person{Name: "王铁蛋", Hobby: []string{"抽中华", "喝牛栏山", "烫杀马特"}}
	p3 := Person{Name: "王铜蛋", Hobby: []string{"抽小熊猫", "喝剑南春", "烫鸡冠头"}}
	people := make([]Person, 0)
	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)

	jsonFile, _ := os.OpenFile("C:/Tencent1803/W2/day008_JSON读写与异常处理/file/八大姨.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err := encoder.Encode(people)
	if err != nil {
		fmt.Println("编码为JSON文件失败,err=", err)
		return
	}
	fmt.Println("编码成功！")
}

/*解码《于谦.json》为结构体*/
func main043() {
	type Person struct {
		Name  string
		Age   int
		Rmb   float64
		Sex   bool
		Hobby []string
	}

	//打开要解码的JSON文件
	jsonFile, _ := os.OpenFile("C:/Tencent1803/W2/day008_JSON读写与异常处理/file/于谦.json", os.O_RDONLY, 0)
	defer jsonFile.Close()

	//创建接收数据的结构体实例指针
	personPtr := new(Person)

	//创建JSON文件的解码器，并将文件数据解码到结构体实例的地址中
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(personPtr)

	//判断解码结果
	if err != nil {
		fmt.Println("解码JSON文件失败，err=", err)
		return
	}

	//从接收数据的地址中读取解码结果
	fmt.Println("解码成功：", *personPtr)
}

/*解码谦哥八大姨JSON文件为map切片*/
func main044() {

	//打开要解码的JSON文件
	jsonFile, _ := os.OpenFile("C:/Tencent1803/W2/day008_JSON读写与异常处理/file/八大姨.json", os.O_RDONLY, 0)
	defer jsonFile.Close()

	//创建接收数据的map切片
	people := make([]map[string]interface{}, 0)

	//创建JSON文件的解码器，并将文件数据解码到map切片的地址中
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(&people)

	//判断解码结果
	if err != nil {
		fmt.Println("解码JSON文件失败，err=", err)
		return
	}

	//从接收数据的地址中读取解码结果
	fmt.Println("解码成功：", people)
}
