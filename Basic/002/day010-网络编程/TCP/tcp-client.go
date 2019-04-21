package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client is Starting!!!!")

	conn, e := net.Dial("tcp", "127.0.0.1:12346")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("Client is Connecting Server")
	//关闭网络连接
	defer conn.Close()

	go func() {
		for {
			time.Sleep(time.Second)
			i, err2 := conn.Write([]byte("Hello from Client !!!\t"))
			if err2 == nil {
				fmt.Println("write len:", i)
			}
		}

	}()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err == nil {
			/*if n != len(buffer) {
				bufferRead := make([]byte, n)
				copy(bufferRead, buffer)
				fmt.Println("收到服务器信息：", string(bufferRead))
			} else {
				fmt.Println("收到服务器信息：", string(buffer))
			}*/
			fmt.Println("收到服务器信息：", string(buffer[:n]))
		}
	}

}
