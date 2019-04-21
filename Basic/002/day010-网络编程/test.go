package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/*
读取目录下的文件
 */
func ReadDirFiles() {
	file, _ := os.Open(`D:\Software\baidu\兄弟连Go基础\Code\day010`)
	names, e := file.Readdirnames(0)
	if e == nil {
		for i := 0; i < len(names); i++ {
			println(names[i])
		}
	}
}

func ReadDir() {
	file, _ := os.Open(`D:\Software\baidu\兄弟连Go基础\Code\day008JSON读写与异常处理`)
	infos, _ := file.Readdir(0)

	for i := 0; i < len(infos); i++ {
		info := infos[i]
		println("文件名:", info.Name())
		println("是否是目录:", info.IsDir())
		println("Mode:", info.Mode())
		println("修改时间:", info.ModTime().Format("2012-10-10 13:13:13"))
		println("***********************************************")
	}

}

func fortest() {
	for true {
		//util.DecodeSLEB128()
		fmt.Println("Hello World")
		time.Sleep(time.Second)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(a))

	s0 := a[:3]
	fmt.Println(s0)

	s1 := a[0:3]
	fmt.Println(s1)

	s2 := a[3:5]
	fmt.Println(s2)

	s3 := a[3:]
	fmt.Println(s3)


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(  text + "\n")
}
