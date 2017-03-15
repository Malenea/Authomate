package main

import (
		"encoding/xml"
		"bytes"
		"fmt"
		"log"
)

// Structures for XML beacons

type Book struct {
	XMLName 	xml.Name	`xml:"book"`
	Id			int			`xml:"id"`
	Title		string 		`xml:"title"`
}

type Library struct {
	XMLName		xml.Name	`xml:"books"`
	Books 		[]Book 		`xml:"book"`
}

type Author struct {
	XMLName		xml.Name	`xml:"author"`
	Name 		string 		`xml:"name"`
	Lib 		Library		`xml:"books"`		
}

type Div struct {
	XMLName		xml.Name	`xml:"GoodreadsResponse"`
	Auth		Author		`xml:"author"`
}

func (b Book) String() string {
	return fmt.Sprintf(" Id : %d - Title : %s\n",
		b.Id, b.Title)
}

// Main function of the XmlParser that creates the url using an id and the author's
// goodreads' key and parses the formatted XML as well as handling errors

func XmlParser(id, key string) {
	var url bytes.Buffer
	url.WriteString("https://www.goodreads.com/author/list.xml?key=")
	url.WriteString(key)
	url.WriteString("&id=")
	url.WriteString(id)

	if XMLdata, err := GetXml(url.String()); err != nil {
		log.Printf("Failed to retrieve XML: %v", err)
	} else {
		var d Div
		xml.Unmarshal(XMLdata, &d)

		fmt.Println("\nAuthor: ", d.Auth.Name, "\n")
		fmt.Println(d.Auth.Lib.Books)
	}
}