package main

import (
	"fmt"
	"net"
	"os/exec"
)

func dealwithconnection(connect net.Conn) {

	fmt.Println("TCP服务器开始处理请求")
	buffer := make([]byte, 1024)
	connect.Read(buffer) //收到消息
	fmt.Println("服务器收到命令" + string(buffer))
	cmd := exec.Command(string(buffer))
	buf, err := cmd.Output() //获取输出
	fmt.Printf("output:%s\n", buf)
	fmt.Printf("error:%s\n", err)

	connect.Write([]byte(buf))
	fmt.Println("TCP服务器处理完请求")

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
		for {
			buffer := make([]byte, 1024)
			n, err := connect.Read(buffer) //收到消息
			if err == nil {

				fmt.Println("TCP服务器开始处理请求")
				fmt.Println(string(buffer[:n]))

				cmd := exec.Command(string(buffer[:n]))
				buf, err := cmd.Output() //获取输出
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("output:%s\n", buf)
				fmt.Printf("error:%s\n", err)

				connect.Write([]byte(buf))
				fmt.Println("TCP服务器处理完请求")
			}

		}

	}

}
