package main

import (
	"fmt"
	"os"
	"strconv"
)

// Constants for Main package

const fileExt = ".html"

// Structs for Main package

type fileDetails struct {
	fileName      string
	numberOfFiles int
}

type details struct {
	users    []string
	mobiles  []string
	channels []string
}

func main() {

	// Declared empty file details to be completed in the swtich statement below
	// If no number of files supplied, go off on a different fork to complete only one
	// Likely also move the swtich statement to a clean file and call form the main

	files := &fileDetails{}
	dt := &details{}

	switch len(os.Args) {
	case 1:
		fmt.Println("no file name provided")
		os.Exit(1)
	case 2:
		fmt.Printf("starting Parse of: %v\n", os.Args[1])
		files.fileName = os.Args[1]
		fmt.Printf("Single file being parsed: %v%v", files.fileName, fileExt)
	case 3:
		fmt.Printf("starting crawl of: %v\n", os.Args[1])
		files.fileName = os.Args[1]
		intConv, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error converting string to int")
			return
		}
		files.numberOfFiles = intConv
		/*
			if err != nil {
				fmt.Println("Error converting input to int")
				return
			}
		*/
		fmt.Printf("File name: %v\nNumber of files to be parsed: %v", files.fileName, files.numberOfFiles)

	/* Concurrency logic - yet to be implemented
	case 4:
		fmt.Printf("starting crawl of: %v\n", os.Args[1])
		website = os.Args[1]
		maxConcurrency, _ = strconv.Atoi(os.Args[2])
		fmt.Printf("Max Concurrency set to : %v\n", maxConcurrency)
		maxPagesSet, _ = strconv.Atoi(os.Args[3])
		fmt.Printf("Max pages set to : %v\n", maxPagesSet)

	*/

	default:
		fmt.Println("Failed to set right parameters")
		return
	}

	files.testPrint()
	hosp := files.createFileName()
	fmt.Println(hosp)

	target := files.createFileName()
	for i := range target {
		resp, err := openHTML(target[i])
		if err != nil {
			fmt.Println("Error getting html string")
			return
		}
		//fmt.Printf("Resp: %v\n", resp)
		dt.GetDetailsFromHTML(resp)
	}

	fmt.Printf("Users: %v", dt.users)
	fmt.Printf("Numbers: %v", dt.mobiles)
	fmt.Printf("Channels: %v", dt.channels)
}
