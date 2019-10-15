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
		fmt.Println(err)
		return
	}
	fmt.Println("client  is  connecting server")
	defer connect.Close()

	//发送消息
	connect.Write([]byte("中国-China"))
	//接收消息
	buffer := make([]byte, 100) //分配内存空间

	n, err := connect.Read(buffer) //读取内容并写入内存空间
	if err == nil {
		fmt.Println("----------------------------读取的信息(来自服务端的消息)----------------------------")
		fmt.Println("[", string(buffer), "]")
		fmt.Println("[", string(buffer[:n]), "]")
		fmt.Println("----------------------------读取的信息----------------------------")
	}
}
