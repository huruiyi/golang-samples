package main

import (
	"fmt"
	"math/rand"
	time2 "time"
)

/*
√ 接口：影视作品，方法：制作、上映；
√ 接口：黄暴产品，方法：刺激你的神经；
√ 接口：东方艺术，继承影视作品和黄暴产品；
√ 封装：封装电影类，属性：名字、公司、主演们，实现影视作品接口；
√ 继承：继承电影类，实现爱情片、科幻片子类，各自覆写制作方法；
√ 多态：影视作品细分：电影、电视剧、网络剧；
√ 类型断言：零点以前看随机电影，零点以后看东方艺术；
*/

/*接口：影视作品，方法：制作、上映*/
type Video interface {
	Make()
	OnShow()
}

/*接口：黄暴产品，方法：刺激你的神经*/
type Violence interface {
	Stimulate()
}

/*接口：东方艺术，继承影视作品和黄暴产品*/
type JapaneseArt interface {
	Video
	Violence
}

/*封装：封装电影类，属性：名字、公司、主演们，实现影视作品接口*/
type Film struct {
	Name    string
	Company string
	Actors  []string
}

func (f *Film) Make() {
	fmt.Println("电影在制作中...")
}
func (f *Film) OnShow() {
	fmt.Println("电影在上映中...")
}

/*多态：影视作品细分：电影、电视剧、网络剧；*/
type TvPlay struct{}

func (tp *TvPlay) Make() {
	fmt.Println("电视剧在制作中...")
}
func (tp *TvPlay) OnShow() {
	fmt.Println("电视剧在上映中...")
}

type NetPlay struct{}

func (tp *NetPlay) Make() {
	fmt.Println("网络剧在制作中...")
}
func (tp *NetPlay) OnShow() {
	fmt.Println("网络剧在上映中...")
}

/*继承电影类，实现爱情片、科幻片子类，各自覆写制作方法*/
type LoveFilm struct {
	Film
}

func (f *LoveFilm) Make() {
	fmt.Println("爱情片在制作中...")
}

type ScienceFilm struct {
	Film
}

func (f *ScienceFilm) Make() {
	fmt.Println("科幻片在制作中...")
}

type LoveActionFilm struct {
	LoveFilm
}

func (laf *LoveActionFilm) Stimulate() {
	fmt.Println("那是十分的刺激")
}

/*实现东方艺术*/

/*类型断言：9-24以前看随机电影，0-9零点以后看东方艺术*/
func main() {
	videos := make([]Video, 0)

	film1 := &Film{Name: "肖申克的救赎", Actors: []string{"蒂姆·罗宾斯", "摩根·弗里曼"}}
	film2 := &LoveFilm{Film{"霸王别姬", "中国公司", []string{"张丰毅", "张国荣"}}}
	film3 := new(ScienceFilm)
	film3.Name = "银翼杀手"

	laf := &LoveActionFilm{LoveFilm{Film{Name: "代码痴汉.avi"}}}

	videos = append(videos, film1)
	videos = append(videos, film2)
	videos = append(videos, film3)
	videos = append(videos, laf)

	//生成随机时间
	time := rand.New(rand.NewSource(time2.Now().UnixNano())).Intn(25)
	fmt.Println("当前时间", time)
	if time > 8 && time < 25 {
		for _, video := range videos {
			switch video.(type) {
			case *Film:
				fmt.Printf("%s可以观看\n", video.(*Film).Name)
			case *LoveFilm:
				fmt.Printf("%s可以观看\n", video.(*LoveFilm).Name)
			case *ScienceFilm:
				fmt.Printf("%s可以观看\n", video.(*ScienceFilm).Name)
			default:
				fmt.Printf("%s不宜观看\n", video.(*LoveActionFilm).Name)
			}
		}
	} else {
		for _, video := range videos {
			if laf, ok := video.(*LoveActionFilm); ok {
				fmt.Printf("愉快地欣赏%s\n", laf.Name)
			}
		}
	}

}
