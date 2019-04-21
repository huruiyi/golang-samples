package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("client  is starting") //客户端开启
	//服务器 127.0.0.1
	connect, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err) //打印错误
		return
	}
	fmt.Println("client  is  connecting server") //链接服务器成功
	defer connect.Close()                        //关闭网络链接

	for {
		//var  cmd string
		//fmt.Scanf("%s",&cmd)  //输入一个命令
		//发送消息
		cmd := "calc"
		connect.Write([]byte(cmd)) //命令传输给服务器
		//接收消息
		buffer := make([]byte, 1024) //分配内存空间
		connect.Read(buffer)         //读取内容并写入内存空间
		fmt.Println(string(buffer))  //显示读取内容
	}

}
