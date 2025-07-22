package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	links    []string
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
		before, found := strings.CutSuffix(os.Args[1], ".html")
		if !found {
			files.fileName = os.Args[1]
		} else {
			files.fileName = before
		}
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
		/* I don't think this part is actually needed - as no Get call is required
		htmlBody, err := GetHTML(resp)
		if err != nil {
			fmt.Println("Error doing HTML parsing")
			fmt.Println(err)
		}
		*/

		dt.GetDetailsFromHTML(resp)
	}
	checkUsers := len(dt.users)
	checkMobiles := len(dt.mobiles)
	checkChannels := len(dt.channels)
	checkLinks := len(dt.links)
	fmt.Printf("Users: %v\n", checkUsers)
	fmt.Printf("Numbers: %v\n", checkMobiles)
	fmt.Printf("Channels: %v\n", checkChannels)
	fmt.Printf("Links: %v\n", checkLinks)
	for i, userN := range dt.users {
		fmt.Printf("User[%v]: %v\n", i, userN)
	}
	for i, mob := range dt.mobiles {
		fmt.Printf("Mobile[%v]: %v\n", i, mob)
	}
	for i, channels := range dt.channels {
		fmt.Printf("User[%v]: %v\n", i, channels)
	}
	for i, userLinks := range dt.links {
		fmt.Printf("User[%v]: %v\n", i, userLinks)
	}
	err := dt.reportHTML()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("===== Report Complete =====")
}
