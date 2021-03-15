package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/codemicro/docuSort/internal/helpers"
	"github.com/codemicro/docuSort/internal/storage"
	"github.com/manifoldco/promptui"
)

var (
	subjects = [...]string{"Computer Science", "Maths", "Physics", "Stats and Mechanics"}
	scanner  = bufio.NewScanner(os.Stdin)
)

const (
	version = "1.2.0"
)

func GetOption(label string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}
	_, selected, err := prompt.Run()
	return selected, err
}

func main() {
	fmt.Println("docuSort save tool v" + version)
	fmt.Println()

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments.")
		return
	}

	document := os.Args[1]

	fmt.Println("Selected document '" + document + "'")

	// Get subject

	fmt.Println("Select subject:")

	for i, v := range subjects {
		fmt.Println(" ", strconv.Itoa(i+1)+":", v)
		_ = os.Mkdir(v, os.ModeDir)
	}

	var selectedSubject string

	for {

		var selectedOption int

		fmt.Print("Option: ")
		scanner.Scan()
		text := scanner.Text()

		var err error
		selectedOption, err = strconv.Atoi(text)

		if err != nil {
			fmt.Println("Not a number.")
			continue
		}

		if selectedOption < 1 || selectedOption > len(subjects) {
			fmt.Println("Out of bounds.")
			continue
		}

		selectedOption -= 1
		selectedSubject = subjects[selectedOption]

		break

	}

	// Get date
	var dateString string

	fmt.Print("Date (blank for today, else yyyy-mm-dd): ")
	scanner.Scan()
	text := scanner.Text()

	if text == "" {
		dateString = time.Now().Format("2006-01-02") // yyyy-mm-dd
	} else {
		dateString = text
	}

	fmt.Println("Topics (enter blank value to finish):")

	var topics []string

	for {

		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {

			if len(topics) == 0 {
				fmt.Println("At least one topic is required")
				continue
			}

			break
		}

		topics = append(topics, text)

	}

	// Get teacher
	fmt.Print("Teacher: ")
	scanner.Scan()
	teacher := scanner.Text()

	// Get type
	documentType, err := GetOption("Type", []string{"Classwork", "Homework", "Assessment"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	existingFiles, err := storage.GetFiles()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	documentComponents := strings.Split(document, ".")
	fileExt := documentComponents[len(documentComponents)-1]

	var topicsForFilename []string
	for _, v := range topics {
		topicsForFilename = append(topicsForFilename, strings.ReplaceAll(strings.ToLower(v), " ", ""))
	}

	newFilename := strings.ReplaceAll(fmt.Sprintf("%s %s.", dateString, strings.Join(topicsForFilename, " "))+fileExt, "/", "")
	newFilename = strings.ReplaceAll(newFilename, "\\", "")

	newFileLocation := filepath.Join(selectedSubject, newFilename)

	thisFile := storage.Document{
		Subject:  selectedSubject,
		Filename: newFileLocation,
		Topics:   topics,
		Teacher:  teacher,
		Type:     documentType,
	}

	existingFiles = append(existingFiles, thisFile)
	fmt.Println(existingFiles)
	err = storage.SaveFiles(existingFiles)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = helpers.CopyFile(document, newFileLocation)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = os.Remove(document)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Saved")

}
