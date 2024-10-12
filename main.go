package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// to ensure the number of arguments
	if len(os.Args) != 3 {
		fmt.Println("USAGE: go run . <inputfile.txt> <outputfile.txt> ")
		return
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]

	// To avoid other files containing other extension and to only have files ending with ".txt"
	if !strings.HasSuffix(inputfile, ".txt") || !strings.HasSuffix(outputfile, ".txt") {
		fmt.Println("error: Both files must have a '.txt' extensions ")
		return
	}
	// os.Stat(inputfile)gets the information for the file
	// os.IsNotExist(er) checks if it does exist
	if _, er := os.Stat(inputfile); os.IsNotExist(er) {
		fmt.Println("error: The input file does not exist")
		return
	}
	// opens the input file and shows an error in case of an error
	content, er := os.Open(inputfile)
	if er != nil {
		fmt.Println("error: error opening the input file", er)
		return
	}
	defer content.Close()
	// checking if it has a content
	infoFile, er := os.Stat(inputfile)
	if er != nil {
		fmt.Println("error: retrieving information", er)
		return
	}
	// In case of an empty file
	// retreiving the size of the file in bytes
	fileSize := infoFile.Size()
	if fileSize == 0 {
		fmt.Println("error: file is empty")
		return
	}
	// this holds the entire content of the file (with the length eq to the file size)
	buffer := make([]byte, fileSize)
	_, er = content.Read(buffer)
	if er != nil {
		fmt.Println("Error reading input file:", er)
		return
	}
	// Preprocess the text to handle punctuation
	processedText := preprocessText(string(buffer))
	// Modify the text based on the other rules
	modifiedText := modifyText(processedText)

	// This will open the file, convert it into byte 
	// Permissions: owner to read and write. And others access to read only
	er = os.WriteFile(outputfile, []byte(modifiedText), 0644)
	if er != nil {
		fmt.Println("error: error writing to output file", er)
		return
	}
}
