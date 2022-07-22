package main

import (
	"fmt"
	"strings"
)

/*
对字符串的处理
- 检索子串
- 格式化
- 比较大小
- 裁剪
- 炸碎
- 拼接
*/

//检索子串
func main051() {

	//fmt.Println(strings.Contains("Hello","el"))//true
	//fmt.Println(strings.Contains("Hello","asshole"))//false
	//fmt.Println(strings.ContainsAny("hello","asshole"))//true
	//fmt.Println(strings.ContainsAny("hello","ass"))//false

	////以字符集【序号】和【字符】形式显示字符
	//fmt.Printf("%U\n",'我')
	//fmt.Printf("%c\n",'我')
	//fmt.Printf("%U\n",0x6211)
	//fmt.Printf("%c\n",0x6211)
	//fmt.Printf("%c\n",0x0068)

	////判断是否包含任意字符
	//fmt.Println(strings.ContainsRune("hello",'h'))//true
	//fmt.Println(strings.ContainsRune("hello",0x0068))//true
	//fmt.Println(strings.ContainsRune("hello",'我'))//false
	//fmt.Println(strings.ContainsRune("额爱你",'我'))//false
	//fmt.Println(strings.ContainsRune("我爱你",'我'))//true
	//fmt.Println(strings.ContainsRune("我爱你",0x6211))//true

	//fmt.Println(strings.Index("fuckoff","fuck"))//0
	//fmt.Println(strings.Index("fuckoff","off"))//4
	//fmt.Println(strings.Index("fuckoff","shit"))//-1
	//fmt.Println(strings.IndexAny("fuckoff","ashole"))//4
	//fmt.Println(strings.IndexRune("fuckoff",0x6211))//-1
	//fmt.Println(strings.IndexRune("fuck我off",0x6211))//4
}

//格式化
func main052() {
	fmt.Println(strings.ToUpper("HeLLo")) //HELLO
	fmt.Println(strings.ToLower("HeLLo")) //hello
	fmt.Println(strings.Title("hello"))   //Hello
}

//比较大小
func main053() {
	fmt.Println(strings.Compare("mfuck", "zshit")) //-1
	fmt.Println(strings.Compare("mfuck", "ashit")) //1
	fmt.Println(strings.Compare("mfuck", "mfuck")) //0
}

//裁剪（只裁头尾，不裁中间）
func main054() {
	////干掉头尾的空格
	//fmt.Println(strings.TrimSpace("   fuck you fuck me   "))
	//fmt.Println(strings.TrimPrefix("合：fuckyoufuckme","合："))
	//fmt.Println(strings.TrimSuffix("合：fuckyoufuckme，ohhh","ohhh"))

	////干掉开头和结尾的f,u,c,k
	//fmt.Println(strings.Trim("fuckyoufuckmeuc","fuck"))//youfuckme
	//fmt.Println(strings.TrimLeft("fuckyoufuckmeuc","fuck"))//youfuckmeuc
	//fmt.Println(strings.TrimRight("fuckyoufuckmeuc","fuck"))//fuckyoufuckme

	//通过函数自定义复杂的裁剪规则
	fmt.Println(strings.TrimFunc("fuckyoufuckmeuc", filter)) //ckyoufuckmeuc
	fmt.Println(strings.TrimFunc("fuckyoufuckmeuu", func(char rune) bool {
		if char == 'f' || char == 'u' {
			return true
		} else {
			return false
		}
	})) //ckyoufuckme
}

/*哪个字符返回true哪个字符就上黑名单——去掉*/
func filter(char rune) bool {
	if char == 'f' || char == 'u' {
		return true
	} else {
		return false
	}
}

//炸碎
func main055() {
	var strs []string
	strs = strings.Split("fuck you fuck me", " ")
	//fmt.Println(strs, len(strs))

	////炸碎为N段，N=-1时有多碎炸多碎
	//strs = strings.SplitN("fuck you fuck me", " ", -1)
	//fmt.Println(len(strs))//4
	//for _,s := range strs{
	//	fmt.Println(s)
	//}

	//将分隔符本身包含进去
	//strs = strings.SplitAfter("fuck,you,fuck,me", ",")
	//fmt.Println(strs, len(strs))
	strs = strings.SplitAfterN("fuck,you,fuck,me", ",", 2)
	fmt.Println(strs, len(strs))
}

//拼接
func main056() {
	bigStr := strings.Join([]string{"fuck", "shit", "asshole"}, "❤")
	fmt.Println(bigStr)
}

/*以字符形式或序号形式输出字符*/
func main057() {
	fmt.Printf("%c\n", 'h')
	fmt.Printf("%c\n", '我')
	fmt.Printf("%U\n", 'h')
	fmt.Printf("%U\n", '我')
}

/*
Go语言中的三种引号
''字符
""字符串
``原样输出
*/
func main() {
	fmt.Println('我')
	fmt.Println("你妹")
	fmt.Println(`
		 1
			2
				3
		你
			大
				爷
	`)
}
