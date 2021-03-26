package main

import "fmt"

/*
封装书的结构体，属性有：书名、价格、作者、状态；
封装读者类，属性有：ReaderID，押金余额；方法有：借书、还书、交罚款；
封装学生类：继承读者类，拓展专业属性、学习方法；
封装老师类：继承读者类，拓展课程属性、教学方法，修改交罚款方法（老师不用交罚款）；
*/
type Book struct {
	Name       string
	Price      float64
	Author     string
	Borrowable bool
}

type Reader struct {
	ReaderID string
	Balance  float64
}

func (r *Reader) BorrowBook(b *Book) {
	fmt.Printf("读者%s借阅了%s\n", r.ReaderID, b.Name)
}
func (r *Reader) ReturnBook(b *Book) {
	fmt.Printf("读者%s归还了%s\n", r.ReaderID, b.Name)
}
func (r *Reader) PayPenalty(amount float64) {
	r.Balance -= amount
	fmt.Printf("读者%s缴纳了%.2f元罚金", r.ReaderID, amount)
}

type Student struct {
	Reader
	Name    string
	Faculty string
}

func (s *Student) Study() {
	fmt.Printf("%s正在学习%s\n", s.Name, s.Faculty)
}

type Teacher struct {
	Reader
	Name   string
	Course string
}

func (t *Teacher) Teach() {
	fmt.Printf("%s正在教授%s\n", t.Name, t.Course)
}
func (r *Teacher) PayPenalty(amount float64) {
	//r.Balance -= amount
	fmt.Println("教师免交罚金")
}

func main() {
	b1 := Book{}
	b1.Name = "三国"
	b1.Author = "罗贯中"
	b1.Borrowable = true
	b1.Price = 12.34
	b2 := Book{"水浒", 34.56, "施耐庵", true}
	b3 := Book{Author: "曹雪芹", Name: "红楼梦"}
	b4 := new(Book)
	b4.Name = "西游记"

	fmt.Println(b1, b2, b3, b4)
	fmt.Printf("%v,%v,%v,%v\n", b1, b2, b3, b4)
	fmt.Printf("%+v,%+v,%+v,%+v\n", b1, b2, b3, b4)
	fmt.Printf("%#v,%#v,%#v,%#v\n", b1, b2, b3, b4)

	reader1 := Reader{"001", 100}
	student := Student{reader1, "张全蛋", "质检"}
	student.Study()
	student.BorrowBook(&b1)
	student.BorrowBook(&b2)
	student.ReturnBook(&b1)
	student.PayPenalty(5)
	fmt.Println("学生余额是", student.Balance)

	teacher := Teacher{Reader{"002", 0}, "王老师", "撩妹秘籍"}
	teacher.Teach()
	teacher.BorrowBook(&b3)
	teacher.BorrowBook(b4)
	teacher.BorrowBook(&b3)
	teacher.PayPenalty(500)
	fmt.Println("老师余额是", teacher.Balance)
}
