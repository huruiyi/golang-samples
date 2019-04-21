package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com/s?wd=%E4%BD%A0%E5%A6%B9")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()               //关闭网页
	bytes, _ := ioutil.ReadAll(resp.Body) //读取全部信息
	fmt.Println(string(bytes))            //打印网页的源代码

}
