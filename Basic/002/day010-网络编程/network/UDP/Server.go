package main

import (
	"fmt"
	"net"
	"os"
)

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when) //显示什么时候发生的什么错误
		os.Exit(1)             //退出
	}
}
func main() {
	fmt.Println("服务器开启")
	udpaddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	ServerHandleError(err, "net.ResolveUDPAddr")

	udpConn, err := net.ListenUDP("udp", udpaddr) //建立UDP链接，得到广口链接
	ServerHandleError(err, "net.ListenUDP")

	//开辟内存缓冲区
	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, _ := udpConn.ReadFromUDP(buffer) //读取UDP的信息
		//打印数据包信息内容
		clientmsg := string(buffer[:n]) //读取客户端信息
		fmt.Printf("收到客户端%v信息%s\n", remoteAddr, clientmsg)
		if clientmsg != "im off" {
			udpConn.WriteToUDP([]byte("已阅:"+clientmsg), remoteAddr)
		} else {
			udpConn.WriteToUDP([]byte("close"), remoteAddr)
		}
	}

	fmt.Println("服务器归天")
}
