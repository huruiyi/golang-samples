package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"io"
	"io/ioutil"
)

/*
√ 简易方式打开一个文件，拿着一顿骚操作，一秒后关闭；
√ 以只读方式打开一个文件，创建其带缓冲的读取器，逐行读取到末尾；
√ 使用ioutil包对《一些逼嗑》进行简易读取
√ 以创写追加或创写覆盖方式打开一个文件，缓冲式写出几行数据，倒干缓冲区后退出；
√ 使用ioutil包进行建议文件写出
√ 使用os包的状态检测结合os.IsNotExist(err)判断文件是否存在
*/

func main021() {
	file, err := os.Open("C:/Tencent1803/W2/day007文件读写/files/一些逼嗑.txt")
	if err == nil {
		fmt.Println("文件打开成功")
	} else {
		fmt.Println("文件打开失败,err=", err)
		return
	}

	//延时（函数返回）关闭打开的文件，释放IO资源
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	fmt.Println("拿着一顿骚操作", file)
	time.Sleep(1 * time.Second)
}

/*
以只读方式打开一个文件，创建其带缓冲的读取器，逐行读取到末尾；
-rwx rw- r--
0764
readable=4,writable=2,executable=1
rw-rw-rw-=0666
*/
func main022() {

	//打开指定文件，已只读模式打开，延用原有权限
	file, err := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/一些逼嗑.txt", os.O_RDONLY, 0)

	//判断文件打开是否成功
	if err!=nil{
		fmt.Println("文件打开失败，err=",err)
		return
	}else{
		fmt.Println("文件打开成功")
	}

	//函数返回前关闭打开的文件
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	//创建该文件的缓冲读取器
	reader := bufio.NewReader(file)

	//循环读取
	for{
		//每次读取到下一处换行符的位置，返回包含换行符在内的字符串
		str, err := reader.ReadString('\n')

		if err == nil{
			//输出读取到缓冲区中的内容
			fmt.Println(str)
		}else{

			//读到文件末尾时退出循环
			if err == io.EOF{
				fmt.Println("已到文件末尾")
				break
			}else{
				//打印错误并退出
				fmt.Println("读取失败，err=",err)
				return
			}
		}
	}
	fmt.Println("文件读取完毕!")

}


/*字符串和字节切片的互化*/
func main023() {
	fmt.Println([]byte("你妹"))
	str := string([]byte{228, 189, 160, 229, 166, 185})
	fmt.Println(str)
}

/*使用ioutil包对《一些逼嗑》进行简易读取*/
func main024() {
	bytes, e := ioutil.ReadFile("C:/Tencent1803/W2/day007文件读写/files/一些逼嗑.txt")
	if e == nil{
		fmt.Println("读取成功，内容=",string(bytes))
	}else{
		fmt.Println("读取失败，err=",e)
	}
}

/*以【创-写-追加】或【创-写-覆盖】模式打开一个文件，缓冲式写出几行数据，倒干缓冲区后退出*/
func main025() {

	//打开指定文件，模式：不存在就创建+只写+覆盖，新创建的文件对所有人可读可写
	file, err := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/二些逼嗑.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err !=nil{
		fmt.Println("文件打开失败，err=",err)
		return
	}

	//函数返回前关闭文件，释放IO资源
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	//创建文件的写出器
	writer := bufio.NewWriter(file)

	//分批次写入字符串
	writer.WriteString("女人的四大梦想\n")
	writer.WriteString("①男人脑壳都坏掉\n")
	writer.WriteString("②天天给我送钞票\n")
	writer.WriteString("③还要排队等我挑\n")
	writer.WriteString("④老娘永远不变老\n")

	//写出一个字节
	writer.WriteByte(123)
	//写出一个字符
	writer.WriteRune('浪')
	//写出一堆字节
	writer.Write([]byte{124,125,126})

	//强制将缓冲区中的内容写出到文件
	writer.Flush()
}

/*使用ioutil包进行简易文件写出*/
func main026() {
	dataStr := `当代四大屌丝
悔创阿里杰克马
	一无所有王健林
		不识妻美刘强东
			普通家庭马化腾
`
	err := ioutil.WriteFile("C:/Tencent1803/W2/day007文件读写/files/三些逼嗑.txt", []byte(dataStr), 0666)
	if err != nil{
		fmt.Println("写出失败，err=",err)
	}else{
		fmt.Println("写出成功！")
	}
}

/*使用os包的状态检测结合os.IsNotExist(err)判断文件是否存在*/
func main027() {
	_, e := os.Stat("C:/Tencent1803/W2/day007文件读写/files/三些2逼嗑.txt")
	if e == nil{
		fmt.Println("文件存在！")
	}else{
		if os.IsNotExist(e){
			fmt.Println("文件不存在")
		}else{
			fmt.Println("发生其他错误，err=",e)
		}
	}
}



