package main

import (
	"fmt"
	"io/ioutil"
)

/*字符串可以视为字符切片进行遍历*/
func main041() {
	for _, c := range "你妹啊" {
		fmt.Printf("%c\n", c)
	}
}

/*对一个文本文件做字符统计，包含字母、数字、空白（'\t',' ','\n','\r')和其它；*/
func main() {
	bytes, _ := ioutil.ReadFile("C:/Tencent1803/W2/day007文件读写/files/temp.txt")
	contentStr := string(bytes)

	var numCount, letterCount,spaceCount,otherCount int
	for _, c := range contentStr {
		if c >= '0' && c <= '9' {
			numCount ++
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			letterCount ++
		} else if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			spaceCount++
		}else{
			otherCount++
		}
	}

	fmt.Println("数字个数为", numCount)
	fmt.Println("字母个数为", letterCount)
	fmt.Println("空白个数为", spaceCount)
	fmt.Println("其它个数为", otherCount)
}
