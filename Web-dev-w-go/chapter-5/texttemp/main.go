package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title       string
	Description string
}

const tmp = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {
	//Create an instance of Note struct
	note1 := Note{"text/templates", "Template generates textual output"}

	//create a new template with a name
	t := template.New("note1")

	//parse some content and generate a template
	t, err := t.Parse(tmp)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note1); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
