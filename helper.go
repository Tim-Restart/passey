package main

import (
	"fmt"
	"strconv"
)

func (files *fileDetails) testPrint() {
	fmt.Printf(`
===================================
File name to be searched: %v%v
Number of files to be searched: %v
===================================
	`, files.fileName, fileExt, files.numberOfFiles)
}

func (files *fileDetails) createFileName() []string {
	// Function to create the file name to parse to the open file command
	// This will check if it is a single file, or a sequential file
	// Logic for this function will likely just deal with single files
	// Second help function to be callled in this one to deal with multiple files?

	// Initalise a slice of strings to load the file names into
	var fileSlice []string

	if files.numberOfFiles != 0 {
		// Logic here for multiple files
		for i := 0; i < files.numberOfFiles; i++ {
			if i == 0 {
				// Edgecase for the first html document, no numbers appended
				fileSlice = append(fileSlice, files.fileName+fileExt)
			} else {
				indexed := i + 1
				n := strconv.Itoa(indexed)
				fileSlice = append(fileSlice, files.fileName+n+fileExt)
			}
		}
		return fileSlice
	} else {
		fileSlice = append(fileSlice, files.fileName+fileExt)
		return fileSlice
	}
}

/*
func (files *fileDetails) openFile(file string)

file, err := os.Open(file) // For read access.
if err != nil {
	log.Fatal(err)
}
*/
