package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codemicro/docuSort/internal/storage"
	"github.com/codemicro/docuSort/internal/templates"
)

const (
	version = "1.2.0"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=../../internal/templates

func main() {
	fmt.Println("docuSort build tool v" + version)
	fmt.Println()

	documents, err := storage.GetFiles()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subjects := storage.GetSubjects(documents)

	var buf bytes.Buffer
	templates.WritePageTemplate(&buf, &templates.HomePage{
		Subjects: subjects,
	})

	ioutil.WriteFile("index.html", buf.Bytes(), 0644)

	for _, subject := range subjects {
		buf.Reset()
		subjectFiles := storage.FilterWhereSubjectIs(subject, documents)
		jr, err := json.Marshal(subjectFiles)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		templates.WritePageTemplate(&buf, &templates.SubjectPage{
			Subjects: subjects,
			Subject:  subject,
			Files:    string(jr),
		})

		_ = os.Mkdir(subject, os.ModeDir)
		ioutil.WriteFile(filepath.Join(subject, "index.html"), buf.Bytes(), 0644)
	}

}
