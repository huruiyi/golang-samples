package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when) //显示什么时候发生的什么错误
		os.Exit(1)             //退出
	}
}

func main() {
	fmt.Println("客户端开启")
	connect, err := net.Dial("udp", "127.0.0.1:8888") //链接服务器，用UDP
	ClientHandleError(err, "net.Dial")                //处理错误
	reader := bufio.NewReader(os.Stdin)               //读取键盘输入
	buffer := make([]byte, 1024)                      //开辟内存空间
	for {
		lineBytes, _, _ := reader.ReadLine() //读取一行
		connect.Write(lineBytes)             //写入一行
		n, _ := connect.Read(buffer)         //读取
		servermsg := string(buffer[:n])      //截取服务器信息
		fmt.Println("服务器", servermsg)        //显示服务器信息
		if servermsg == "close" {
			connect.Write([]byte("你妹的服务器，大爷我不稀罕你"))
			break
		}
	}
	fmt.Println("客户端结束")

}
