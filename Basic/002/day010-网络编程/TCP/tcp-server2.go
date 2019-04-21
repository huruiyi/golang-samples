package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func read2(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == '\t' {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func handCon(conn net.Conn) {

	defer conn.Close()
	for {
		s, err2 := read2(conn)
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("The connection is closed by another side.")
			} else {
				fmt.Println("Read Error: %s", err2)
			}
		}
		fmt.Println(conn.RemoteAddr(), ":", s)
		conn.Write([]byte("来自服务器：收到消息"))
	}

}

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
			fmt.Println("客户端连接上了:", conn.RemoteAddr())
			go handCon(conn)
		}

	}
}
