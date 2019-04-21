package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好baby"))
	})
	http.HandleFunc("/bitcoin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好baby  bitcoin"))
	})
	http.HandleFunc("/ethereum", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好baby  ethereum"))
	})
	http.HandleFunc("/eos", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好baby  eos"))
		writer.Write([]byte("请求方法：" + request.Method))
		writer.Write([]byte("请求内容长度:" + strconv.Itoa(int(request.ContentLength))))
		writer.Write([]byte("请求主机：" + request.Host))
		writer.Write([]byte("请求协议" + request.Proto))
		writer.Write([]byte("请求远程地址" + request.RemoteAddr))
		writer.Write([]byte("请求路由" + request.RequestURI))
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
