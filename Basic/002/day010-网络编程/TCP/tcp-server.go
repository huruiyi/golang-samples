package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is Starting")

	listener, e := net.Listen("tcp", "127.0.0.1:12346")
	if e != nil {
		fmt.Println("error", e)
		return
	}
	defer listener.Close()

	fmt.Println("Server is starting listen tcp 127.0.0.1:12346")
	for true {

		conn, err := listener.Accept()
		if err == nil {
			buffer := make([]byte, 1024)
			n, err2 := conn.Read(buffer)

			if err2 == nil {
			/*	fmt.Println("读取长度：", n, "buffer.len=", len(buffer))
				if n != len(buffer) {
					bufferRead := make([]byte, n)
					copy(bufferRead, buffer)
					fmt.Println("客户端消息：", string(bufferRead))
				} else {
					fmt.Println("客户端消息：", string(buffer))
				}*/
				fmt.Println("客户端消息：", string(buffer[:n]))
				conn.Write([]byte("来自服务器：收到消息"))
			}
		}
	}
}
