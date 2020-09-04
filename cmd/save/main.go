package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/codemicro/docuSort/internal/storage"
)

var (
	subjects = [...]string{"Computer Science", "Maths", "Physics"}
	scanner  = bufio.NewScanner(os.Stdin)
)

const (
	version = "0.0.0a"
)

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	fmt.Println("docuSort v" + version)
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

	for {

		fmt.Print("Date (blank for today, else ddmmyy): ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {
			dateString = time.Now().Format("020106") // ddmmyy format
			break
		}

		re := regexp.MustCompile(`^\d{8}$`)
		if len(re.FindAllString(text, -1)) == 0 {
			fmt.Println("Incorrect format")
			continue
		}

		dateString = text
		break

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

	existingFiles, err := storage.GetFiles()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subjectFileCount := storage.CountWhereSubjectIs(selectedSubject, existingFiles)

	fmt.Println(subjectFileCount)

	documentComponents := strings.Split(document, ".")
	fileExt := documentComponents[len(documentComponents)-1]

	newFileNumber := subjectFileCount + 1

	var topicsForFilename []string
	for _, v := range topics {
		topicsForFilename = append(topicsForFilename, strings.ReplaceAll(strings.ToLower(v), " ", ""))
	}

	newFileLocation := filepath.Join(selectedSubject, fmt.Sprintf("%04d %s %s.", newFileNumber, dateString, strings.Join(topicsForFilename, " "))+fileExt)

	thisFile := storage.Document{
		Subject:  selectedSubject,
		Filename: newFileLocation,
		Topics:   topics,
	}

	existingFiles = append(existingFiles, thisFile)
	fmt.Println(existingFiles)
	err = storage.SaveFiles(existingFiles)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = copyFile(document, newFileLocation)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Saved")

}