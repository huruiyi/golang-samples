package mylib

import "testing"

func TestEncodePerson2JsonFile(t *testing.T) {
	filename := "C:/Tencent1803/W2/day009_程序测试+反射/demos/file/目标人员.json"

	p0 := Person{"于谦", 50, 123.45, true, []string{"抽烟", "喝酒", "烫头"}}
	ok := EncodePerson2JsonFile(&p0, filename)
	if ok {
		p1, _ := DecodeJsonFile2Person(filename)
		if p1.Name != p0.Name || p1.Age != p0.Age {
			//t.Log("编码的数据和原来不一样！")
			//t.FailNow()
			t.Fatal("编码的数据和原来不一样！")
		} else {
			t.Log("TestEncodePerson2JsonFile 测试成功")
		}
	}
}

func TestDecodeJsonFile2Person(t *testing.T) {
	filename := "C:/Tencent1803/W2/day009_程序测试+反射/demos/file/原始人员.json"
	p0 := Person{"于谦", 50, 123.45, true, []string{"抽烟", "喝酒", "烫头"}}
	ok := EncodePerson2JsonFile(&p0, filename)
	if ok {
		p1, _ := DecodeJsonFile2Person(filename)
		if p1.Name != p0.Name || p1.Age != p0.Age {
			//t.Log("编码的数据和原来不一样！")
			//t.FailNow()
			t.Fatal("编码的数据和原来不一样！")
		} else {
			t.Log("TestEncodePerson2JsonFile 测试成功")
		}
	}
}
