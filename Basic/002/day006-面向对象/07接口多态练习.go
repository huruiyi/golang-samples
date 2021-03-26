package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
定义动物接口：活着,死
定义动物实现接口：鸟、鱼、野兽（跑、捕食）
继承野兽类：实现老虎，实现人
业务场景：工作日所有动物都活着、周末人出来捕食，野兽逃跑，其它动物死光光
*/
type Animal interface {
	Live()
	GoDie()
}

type Bird struct{}

func (b *Bird) Live() {
	fmt.Println("一只鸟儿在唱歌")
}
func (b *Bird) GoDie() {
	fmt.Println("一只鸟儿死翘翘")
}

type Fish struct{}

func (b *Fish) Live() {
	fmt.Println("一只鱼儿在游泳")
}
func (b *Fish) GoDie() {
	fmt.Println("一只鱼儿死翘翘")
}

type Beast struct{}

func (b *Beast) Live() {
	fmt.Println("一只野兽在嬉戏")
}
func (b *Beast) GoDie() {
	fmt.Println("一只野兽死翘翘")
}
func (b *Beast) Run() {
	fmt.Println("一只野兽在奔跑")
}
func (b *Beast) Hunt() {
	fmt.Println("一只野兽在捕食")
}

type Tiger struct {
	Beast
}
type Human struct {
	Beast
	Name string
}

func (h *Human) Hunt() {
	fmt.Printf("%s在捕食\n", h.Name)
}
func (h *Human) Live() {
	fmt.Printf("%s正在拓展他的食谱\n", h.Name)
}

func main() {
	bird := Bird{}
	fish := Fish{}
	tiger := Tiger{Beast{}}
	humanPtr := new(Human)
	humanPtr.Name = "南粤食神"

	//加入地球大家庭
	animals := make([]Animal, 0)
	animals = append(animals, &bird)
	animals = append(animals, &fish)
	animals = append(animals, &tiger)
	animals = append(animals, humanPtr)

	//生成星期日
	weekday := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(7)
	fmt.Println("今天星期", weekday)
	if weekday > 0 && weekday < 6 {
		for _, animal := range animals {
			animal.Live()
		}
	} else {
		for _, animal := range animals {

			switch animal.(type) {
			case *Human:
				animal.(*Human).Hunt()
			case *Tiger:
				animal.(*Tiger).Run()
			default:
				animal.GoDie()
			}

			/*			if human, ok := animal.(*Human); ok {
							human.Hunt()
						} else if tiger, ok := animal.(*Tiger); ok {
							tiger.Run()
						} else {
							animal.GoDie()
						}*/

		}
	}

}
