package main

import (
	"io/ioutil"
	"fmt"
	"io"
	"os"
	"bufio"
)

/*
√ 使用ioutil包做一个傻瓜式拷贝
√ 使用io.Copy进行文件拷贝
√ 使用缓冲1K的缓冲区配合缓冲读写器进行图片拷贝
*/

/*使用ioutil包做一个傻瓜式拷贝*/
func main031() {
	srcBytes, _ := ioutil.ReadFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2.jpg")
	err := ioutil.WriteFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2333.jpg", srcBytes, 0666)
	if err != nil {
		fmt.Println("拷贝失败，err=", err)
		return
	}
	fmt.Println("done!")
}

/*使用io.Copy进行文件拷贝*/
func main032() {
	//打开源文件和目标文件
	srcFile, _ := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2.jpg", os.O_RDONLY, 0)
	dstFile, _ := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2333.jpg", os.O_CREATE|os.O_WRONLY, 0666)

	//函数返回前关闭两个文件
	defer func() {
		srcFile.Close()
		dstFile.Close()
		fmt.Println("文件已关闭")
	}()

	//将源文件的数据拷贝到目标文件
	written, err := io.Copy(dstFile, srcFile)

	//判断拷贝是否成功
	if err != nil {
		fmt.Println("拷贝失败，err=", err)
		return
	}
	fmt.Println("拷贝成功，字节数=", written)
}

/*使用缓冲1K的缓冲区配合缓冲读写器进行大文件拷贝*/
func main033() {
	//打开源文件和目标文件
	srcFile, _ := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2.jpg", os.O_RDONLY, 0)
	dstFile, _ := os.OpenFile("C:/Tencent1803/W2/day007文件读写/files/fuckoff2333.jpg", os.O_CREATE|os.O_WRONLY, 0666)

	//延时关闭文件
	defer func() {
		srcFile.Close()
		dstFile.Close()
		fmt.Println("文件已关闭")
	}()

	//创建源文件的缓冲读取器
	reader := bufio.NewReader(srcFile)
	//创建目标文件的缓冲写出器
	writer := bufio.NewWriter(dstFile)

	//创建小水桶（缓冲区）
	buffer := make([]byte, 1024)

	//一桶一桶地读入源文件数据，直到文件末尾
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("已到文件末尾")
				//倒干最后一个缓冲区的数据
				writer.Flush()
				break
			} else {
				fmt.Println("读取数据失败，err=", err)
				return
			}
		}

		//将读到的数据一桶一桶地写出到目标文件
		_, err = writer.Write(buffer)
		if err != nil {
			fmt.Println("写出数据失败，err=", err)
			return
		}
	}
	fmt.Println("拷贝成功！")
}
