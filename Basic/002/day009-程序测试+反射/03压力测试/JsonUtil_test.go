package mylib

import "testing"

func BenchmarkEncodePerson2JsonFile(b *testing.B) {
	filename := "C:/Tencent1803/W2/day009_程序测试+反射/demos/file/目标人员.json"
	p0 := Person{"于谦", 50, 123.45, true, []string{"抽烟", "喝酒", "烫头"}}

	b.Log("BenchmarkEncodePerson2JsonFile 测试开始")
	//汇报内存分配情况
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EncodePerson2JsonFile(&p0, filename)
	}
}

func BenchmarkDecodeJsonFile2Person(b *testing.B) {
	b.Log("BenchmarkDecodeJsonFile2Person 测试开始")
	filename := "C:/Tencent1803/W2/day009_程序测试+反射/demos/file/原始人员.json"

	//汇报内存分配情况
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DecodeJsonFile2Person(filename)
	}
}
