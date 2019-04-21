package mylib

/*
		t.Log("错误日志")
		t.Logf("测试失败，期望%d,实际得到%d\n",55,sum)

		//宣告测试失败，但还会继续执行下去
		t.Fail()
		//宣告测试失败，并立即终止测试
		t.FailNow()

		//t.Log + t.Fail()
		t.Error("错误日志")
		//t.Logf()+t.Fail()
		t.Errorf("测试失败，期望%d,实际得到%d\n",55,sum)

		//t.Log + t.FailNow()
		t.Fatal("错误日志")
		//t.Logf()+t.FailNow()
		t.Fatalf("测试失败，期望%d,实际得到%d\n",55,sum)
*/

import "testing"

func TestGetSum(t *testing.T) {
	sum := GetSum(10)
	if sum != 55{
		t.Fatalf("测试失败，期望%d,实际得到%d\n",55,sum)
	}
	t.Log("测试成功！")
}

func TestGetSumRecursively(t *testing.T) {
	sum := GetSumRecursively(10)
	if sum != 55{
		t.Fatalf("测试失败，期望%d,实际得到%d\n",55,sum)
	}
	t.Log("测试成功！")
}
