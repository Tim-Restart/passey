package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// This needs to pass to the GetDetails function below

func openHTML(file string) (string, error) {

	html, err := os.ReadFile(file) // Reads the whole file and loads to mem
	if err != nil {
		fmt.Println("Error opening file")
		return "", err
	}

	return string(html), nil

}

// adding phone locating and email?

func (dt *details) emailPhone(n *html.Node, emailRegex, phoneRegex *regexp.Regexp) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			emails := emailRegex.FindAllString(text, -1)
			phones := phoneRegex.FindAllString(text, -1)
			for _, e := range emails {
				//fmt.Println("Found Email:", e)
				cfg.addToEmail(e)
			}
			for _, p := range phones {
				//fmt.Println("Found phone:", p)
				cfg.addToPhone(p)
			}
		}
	}
	// Added check for html.ElementNodes also
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			// Check if attribute value contains emails/phones (like href="mailto:...")
			attrValue := strings.TrimSpace(attr.Val)
			emails := emailRegex.FindAllString(attrValue, -1)
			phones := phoneRegex.FindAllString(attrValue, -1)
			for _, e := range emails {

				cfg.addToEmail(e)
			}
			for _, p := range phones {
				cfg.addToPhone(p)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cfg.emailPhone(c, emailRegex, phoneRegex)
	}
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
				// First check is for mobile numbers and usernames

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

					if strings.Contains(n.FirstChild.Data, "https://t.me") {

						dt.channels = append(dt.channels, n.FirstChild.Data)
					} else if strings.Contains(n.FirstChild.Data, "https://") {
						dt.links = append(dt.links, n.FirstChild.Data)
					} else {
						break
					}
				}

			}
		} else if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// Check if a.Val has a suffix here
					if strings.HasPrefix(a.Val, "http://t.me") {
						if slices.Contains(dt.channels, a.Val) {
							break
						} else {
							dt.channels = append(dt.channels, strings.TrimSpace(a.Val))
						}
					} else if strings.HasPrefix(a.Val, "http://") {
						if slices.Contains(dt.links, a.Val) {
							break
						} else {
							dt.links = append(dt.links, strings.TrimSpace(a.Val))
						}
					}
				}
			}
		}
	}
	return nil
}
