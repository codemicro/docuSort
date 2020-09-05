package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/codemicro/docuSort/internal/templates"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=../../internal/templates

func main() {
	fmt.Println("AAAAAA")

	var buf bytes.Buffer
	templates.WritePageTemplate(&buf, &templates.MainPage{
		Subjects: []string{"Computer Science", "Maths", "Physics"},
	})

	fmt.Printf("buf=\n%s", buf.Bytes())

	ioutil.WriteFile("index.html", buf.Bytes(), 0644)

}
