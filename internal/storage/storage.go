package storage

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var (
	fileName = "documents.json"
)

func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return true
		}
	}
	return false
}

type Document struct {
	Subject  string
	Filename string
	Topics   []string
	Teacher string
	Type string
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
	j, err := json.MarshalIndent(in, "", "  ")
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

func FilterWhereSubjectIs(subject string, documents []Document) []Document {
	var newSlice []Document
	for _, v := range documents {
		if strings.ToLower(subject) == strings.ToLower(v.Subject) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func GetSubjects(documents []Document) []string {
	var subjects []string
	for _, v := range documents {
		if !isStringInSlice(v.Subject, subjects) {
			subjects = append(subjects, v.Subject)
		}
	}
	return subjects
}
