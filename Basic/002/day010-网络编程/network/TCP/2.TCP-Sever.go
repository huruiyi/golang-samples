package main

import (
	"fmt"
	"net"
)

func dealwithconnection(connect net.Conn) {
	fmt.Println("TCP服务器开始处理请求")
	buffer := make([]byte, 100)
	n, err := connect.Read(buffer) //收到消息
	if (err == nil) {
		fmt.Println("读取字节数:", n)
		fmt.Println("----------------------------读取的信息(来自客户端的消息)----------------------------")
		fmt.Println("[", string(buffer), "]")
		fmt.Println("[", string(buffer[:n]), "]")
		fmt.Println("----------------------------读取的信息----------------------------")
	}

	connect.Write([]byte ("收到消息-来自服务器"))
}

func main() {

	fmt.Println("Server  is starting")
	server, err := net.Listen("tcp", "127.0.0.1:12345") //创建服务器
	if err != nil {
		fmt.Println(err) //处理错误
		return
	}
	defer server.Close() //关闭服务器
	fmt.Println("Server  is starting  listen tcp 127.0.0.1:12345")
	for {
		connect, err := server.Accept() //服务器接收客户端请求
		if err != nil {
			fmt.Println(err) //处理错误
			return           //处理错误
		}
		dealwithconnection(connect) //处理客户端请求
	}
}
