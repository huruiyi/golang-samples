package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	fmt.Println("filename:", filename)
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p1 := &Page{Title: "sample", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("sample")
	fmt.Println(string(p2.Body))

	//http://localhost:9000/monkeys
	http.HandleFunc("/", handler)

	//http://127.0.0.1:9000/view/sample
	//根路径需要存在 sample.txt
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
