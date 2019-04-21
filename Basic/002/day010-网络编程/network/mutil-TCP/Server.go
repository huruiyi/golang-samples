package main

import (
	"fmt"
	"net"
	"os"
)

//处理错误
func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when) //显示错误
		os.Exit(1)             //退出程序
	}
}

func ChatWith(conn net.Conn) {
	buffer := make([]byte, 1024) //开辟内存
	for {
		n, err := conn.Read(buffer)         //读取消息
		ServerHandleError(err, "conn.Read") //处理错误
		clientmsg := string(buffer[0:n])    //转化为字符串输出
		fmt.Printf("收到消息%v,%s", conn.RemoteAddr(), clientmsg)
		if clientmsg != "im off" {
			conn.Write([]byte("已阅:" + clientmsg))
		} else {
			conn.Write([]byte("bye")) //退出
			break
		}
	}
	conn.Close()
	fmt.Printf("客户端%s已经断开\n", conn.RemoteAddr())
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	ServerHandleError(err, "net.Listen") //处理错误
	for {
		conn, err := listener.Accept()            //处理客户端链接
		ServerHandleError(err, "listener.Accept") //处理错误
		go ChatWith(conn)                         //协程并发处理
	}
}
