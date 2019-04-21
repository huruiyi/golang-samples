package mylib

import "testing"

func BenchmarkGetSum(b *testing.B) {
	b.Log("BenchmarkGetSum 测试开始")
	//汇报内存分配情况
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSum(10)
	}
}

func BenchmarkGetSumRecursively(b *testing.B) {
	b.Log("BenchmarkGetSumRecursively 测试开始")
	//汇报内存分配情况
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSumRecursively(10)
	}
}
