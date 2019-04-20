package main

import "fmt"
import "net/http"
import "io/ioutil"

func httpGet(url string) string {

	//发送一个url
	resp, err := http.Get(url)
	if err != nil {
		//返回失败
		fmt.Println("http get error")
		return ""
	}

	//resp 里面存放的就是从url返回的数据
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//返回失败
		fmt.Println("readall error")
		return ""
	}

	content := string(data)

	resp.Body.Close()
	return content
}

func main() {

	content := httpGet("http://www.baidu.com/index.html")

	fmt.Println("从百度获取的数据 ", content)

}
