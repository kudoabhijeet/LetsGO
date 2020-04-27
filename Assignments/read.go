package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func getInputFileNameFromPromt() (inputFilePath string, err error) {

	inputFilePath = ""

	for {
		fmt.Println("\nPlease enter the full path for the input file to read: ")
		fmt.Println("or enter 'END' or press <CTRL + C> to exit program")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputFilePath = scanner.Text()
			break
		}

		if err := scanner.Err(); err != nil {
			return "", err
		}

		inputFilePath = strings.TrimSpace(inputFilePath)
		if inputFilePath == "" {
			fmt.Println("Nothing entered. Please try again")
			continue
		} else {
			return inputFilePath, nil
		}
	}

}
func fileExists(name string) bool {
	fileInfo, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !fileInfo.IsDir()
}

func processFile(pathToFile string, personsSlice *[]person) error {
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aPerson := getPerson(scanner.Text())
		*personsSlice = append(*personsSlice, aPerson)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func getPerson(fileLine string) person {

	fileLine = strings.TrimSpace(fileLine)
	lineComponents := strings.Split(fileLine, " ")

	fname := ""
	lname := ""
	if len(lineComponents) >= 2 {
		fname = lineComponents[0]
		lname = lineComponents[1]
	} else {
		fname = lineComponents[0]
		lname = "(N/A)"
	}

	if len(fname) > 20 {
		fname = substr(fname, 0, 20)
	} else if len(fname) == 0 {
		fname = "(N/A)"
	}

	if len(lname) > 20 {
		lname = substr(lname, 0, 20)
	} else if len(lname) == 0 {
		lname = "(N/A)"
	}

	return person{fname, lname}
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func main() {

	// The "empty person slice of max. initial capacity 100 "
	personsSlice := make([]person, 0, 100)

	for {

		inputFilePath, err := getInputFileNameFromPromt()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error condition raised. The program will exit.")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if inputFilePath == "END" {
			fmt.Println("\nGoodbye !!!")
			os.Exit(0)
		}

		if fileExists(inputFilePath) {
			fmt.Printf("\nProcessing file[%s] ...\n", inputFilePath)
			err := processFile(inputFilePath, &personsSlice)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error condition raised. The program will exit.")
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Println("\nIterating through the resulting slice of persons:")
				for n, p := range personsSlice {
					fmt.Printf("Entry #[%d] First Name[%s] Last Name[%s]\n", n, p.fname, p.lname)
				}
				os.Exit(0)
			}
		} else {
			fmt.Printf("File [%s] Doesn't exists.\n", inputFilePath)
			continue
		}
	}

}
