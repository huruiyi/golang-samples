package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 01; i <= 84; i++ {
		dirName := fmt.Sprintf("%d", i)
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			fmt.Printf("创建目录 %s 失败: %v\n", dirName, err)
			return
		}
		fmt.Printf("成功创建目录: %s\n", dirName)
	}
}
