package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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

// Example GetHTML
/*
func GetHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 399 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return "", err
	}
	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		err = fmt.Errorf("Header not text/html")
		return "", err
	}
	if err != nil {
		return "", err
	}
	return string(body), nil

}

*/

// New reader function for the HTML

func (dt *details) GetDetailsFromHTML(htmlBody string) error {

	// get the URL's from the HTML here

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

	//baseURL, err := url.Parse(rawBaseURL)
	//if err != nil {
	//	fmt.Println("Error parsing baseURL string")
	//	return nil, err
	//}

	htmmlReader := strings.NewReader(htmlBody)
	nodeTree, err := html.Parse(htmmlReader)
	if err != nil {
		fmt.Println("Error parsing HTML data to nodes")
		log.Fatal(err)
	}

	for n := range nodeTree.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "from_name" {
					// Check if a.Val has a suffix here
					if strings.HasPrefix(a.Val, "+61") {
						dt.mobiles = append(dt.mobiles, strings.TrimSpace(a.Val))
						break
					} else {
						dt.users = append(dt.users, a.Val)
						break
					}
				} else if a.Key == "text" {
					// checks for a channel, at the moment it will append the whole text string
					if strings.Contains(a.Val, "https://t.me/") {
						dt.channels = append(dt.channels, a.Val)
					}
				}
			}
		}
	}

	// Probably better the change this return to a struct
	// Then it will only be one return
	return nil

}
