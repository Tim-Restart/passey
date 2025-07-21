package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

// create a empty details struct to pass around for use

// This needs to pass to the GetDetails function below

func openHTML(file string) (string, error) {

	html, err := os.ReadFile(file) // Reads the whole file and loads to mem
	if err != nil {
		fmt.Println("Error opening file")
		return "", err
	}

	return string(html), nil

}

// New reader function for the HTML

func (dt *details) GetDetailsFromHTML(htmlBody string) error {

	// parse the URL data to break it down into nodes
	// Nodes are a type as per below:
	/*
			type Node struct {
			Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

			Type      NodeType
			DataAtom  atom.Atom
			Data      string
			Namespace string
			Attr      []Attribute
		}

	*/

	htmmlReader := strings.NewReader(htmlBody)
	nodeTree, err := html.Parse(htmmlReader)
	if err != nil {
		fmt.Println("Error parsing HTML data to nodes")
		log.Fatal(err)
	}
	//fmt.Println("Entering soughting mode")
	for n := range nodeTree.Descendants() {
		if n.Type == html.ElementNode && n.Data == "div" {

			//fmt.Printf("n.Data: %v\n", n.Data)
			for _, a := range n.Attr {
				// First check is for mobile numbers only
				// This could be converted to usernames also
				if a.Key == "class" && a.Val == "from_name" {
					if strings.Contains(n.FirstChild.Data, "+61") {
						//fmt.Printf("Mobiles: %v\n", n.FirstChild.Data)
						// Need to add a check for exists
						if slices.Contains(dt.mobiles, n.FirstChild.Data) {
							//fmt.Println("Slice contains number!")
							break
						} else {
							dt.mobiles = append(dt.mobiles, n.FirstChild.Data)
						}
						// Checks for username - deleted it just ignores
					} else if !strings.Contains(n.FirstChild.Data, "Deleted Account") {

						//fmt.Printf("Entered not deleted check\n")
						// Checks to see if it exists
						if slices.Contains(dt.users, n.FirstChild.Data) {

							break
						} else {
							dt.users = append(dt.users, n.FirstChild.Data)
							break
						}
						// End user and mobile check
					}

				} else if a.Key == "class" && a.Val == "text" {
					//fmt.Println("++++++========== Text contains user ===========")
					if strings.Contains(n.FirstChild.Data, "https://t.me") {
						//fmt.Println("++++++========== append channels ===========")
						dt.channels = append(dt.channels, n.FirstChild.Data)
					} else if strings.Contains(n.FirstChild.Data, "https://") {
						dt.links = append(dt.links, n.FirstChild.Data)
					} else {
						break
					}
				}

			}
		}

	}
	return nil
}
