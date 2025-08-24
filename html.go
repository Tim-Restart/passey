package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"regex"

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



// New reader function for the HTML

func (dt *details) GetDetailsFromHTML(htmlBody string) error {

	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	phoneRegex := regexp.MustCompile(`(?i)(?:\+61|61|\(\+61\))?[\s\-\.]*?(?:\(0?[2-478]\)[\s\-\.]*?\d{4}[\s\-\.]*?\d{4}|0?4\d{2,3}[\s\-\.]*?\d{3}[\s\-\.]*?\d{3})`)

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

	for n := range nodeTree.Descendants() {
		if n.Type == html.ElementNode && n.Data == "div" {


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
		}
		// Email and phones regex checks
		// Uses the regex expressions stated at the start
		// Checks html.Text nodes for the presence of these regex

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

					dt.email(e)
				}
				for _, p := range phones {
					dt.mobiles(p)
				}
			}
		}
	/* Recursion for email/phone call - not sure if needed as using decendents
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cfg.emailPhone(c, emailRegex, phoneRegex)
	}
		*/


	}
	return nil
}


