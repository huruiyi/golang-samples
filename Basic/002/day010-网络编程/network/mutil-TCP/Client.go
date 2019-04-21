package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//处理错误
func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when) //显示错误
		os.Exit(1)             //退出程序
	}
}

func main() {
	fmt.Println("client is start")
	conn, err := net.Dial("tcp", "127.0.0.1:8888") //链接服务器
	ClientHandleError(err, "net.Dial")             //处理错误

	buffer := make([]byte, 1024)        //开辟内存
	reader := bufio.NewReader(os.Stdin) //从命令行读取输入
	for {
		lineBytes, _, _ := reader.ReadLine() //读取一行
		conn.Write(lineBytes)                //给服务器发送消息
		n, _ := conn.Read(buffer)            //读取服务器消息，存入消息缓冲区
		servermsg := string(buffer[0:n])     //截取字符串
		fmt.Println("服务器", servermsg)

		//处理服务器
		if servermsg == "bye" {
			break //收到消息bye,退出死循环
		}

	}

	fmt.Println("client is end")
}
