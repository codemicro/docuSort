package storage

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var (
	fileName = "documents.json"
)

type Document struct {
	Subject  string
	Filename string
	Topics   []string
}

func GetFiles() ([]Document, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []Document{}, nil // Assume it means file does not exist
	}

	var returnValue []Document
	err = json.Unmarshal(content, &returnValue)
	if err != nil {
		return []Document{}, err
	}

	return returnValue, nil

}

func SaveFiles(in []Document) error {
	j, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, j, 0644)
}

func CountWhereSubjectIs(subject string, documents []Document) int32 {
	var count int32
	for _, v := range documents {
		if strings.ToLower(subject) == strings.ToLower(v.Subject) {
			count += 1
		}
	}
	return count
}
