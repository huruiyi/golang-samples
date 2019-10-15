package main

import (
	"fmt"
	"os/exec"
)

func main1() {
	cmd := exec.Command("calc")
	buf, err := cmd.Output() //获取输出
	fmt.Printf("output:%s\n", buf)
	fmt.Printf("error:%s\n", err)
}

func main() {
	cmdstr := "tasklist"
	cmdbytes := []byte(cmdstr) //字符串转化字节
	//newcmdstr:=string(cmdbytes) //字节转化字符串
	cmd := exec.Command(string(cmdbytes))
	buf, err := cmd.Output() //获取输出
	fmt.Printf("output:%s\n", buf)
	fmt.Printf("error:%s\n", err)
}
