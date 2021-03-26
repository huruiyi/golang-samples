package mylib

import "testing"

func TestGetSum2(t *testing.T) {
	answerMap := make(map[int]int)
	answerMap[10] = 55
	answerMap[5] = 15

	for n, answer := range answerMap {
		sum := GetSum(n)
		if sum != answer {
			t.Logf("测试失败，期望%d，实际得到%d\n", answer, sum)
			t.FailNow()
		}
	}
}
