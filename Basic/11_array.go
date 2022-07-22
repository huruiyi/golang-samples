package main

import (
	"fmt"
	"time"
)

func demo1() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	fmt.Print(arr)
	slice = append(slice, 6, 7)
	fmt.Print(arr)
	fmt.Print(slice)
}

func demo2() {
	var arr = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s1 := arr[2:6]
	s2 := s1[1:3]
	fmt.Println(len(s1), cap(s1), string(s1))
	fmt.Println(len(s2), cap(s2), string(s2)+"\n")
	s2 = append(s2, 'l')
	fmt.Println(len(s1), cap(s1), string(s1))
	fmt.Println(len(s2), cap(s2), string(s2)+"\n")
	s3 := make([]byte, 3, 3)
	copy(s3, s2)
	fmt.Println(len(s1), cap(s1), string(s1))
	fmt.Println(len(s2), cap(s2), string(s2))
	fmt.Println(len(s3), cap(s3), string(s3))
}

func demo3() {
	var favorites []string
	favorites = make([]string, 3)
	favorites[0] = "github.com/gorilla/mux"
	favorites[1] = "github.com/codegangsta/negroni"
	favorites[2] = "gopkg.in/mgo.v2"
	for e := range favorites {
		fmt.Println(e, favorites[e])
	}

}

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

func SortSlice(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] > s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println(s)
}

func PrintMatrix() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			fmt.Print("大爷\t\t")
		}
		fmt.Println()
	}
}

func main() {

	var note1 Note
	note1.Title = "流浪地球"
	note1.Description = "科幻电影"
	note1.CreatedOn = time.Now()

	var note2 Note
	note2.Title = "反贪风暴4"
	note2.Description = "卧底牢狱 闯牢打虎"
	note2.CreatedOn = time.Now()

	notes := make(map[string]Note)
	notes["film1"] = note1
	notes["film2"] = note2

	fmt.Println(notes)

	PrintMatrix()

	SortSlice([]int{5, 7, 2, 1, 4, 9, 3, 6, 8, 0})

	demo1()
	demo2()
	demo3()
}
