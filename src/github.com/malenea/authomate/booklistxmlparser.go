package main

import (
		"encoding/xml"
		"fmt"
		"log"
)

// Structures for XML beacons

type Book struct {
	XMLName 	xml.Name	`xml:"book"`
	Id			int			`xml:"id"`
	Title		string 		`xml:"title"`
	NakedTitle	string 		`xml:"title_without_series"`
}

type Library struct {
	XMLName		xml.Name	`xml:"books"`
	Books 		[]Book 		`xml:"book"`
}

type Author struct {
	XMLName		xml.Name	`xml:"author"`
	Id 			int			`xml:"id"`
	Name 		string 		`xml:"name"`
	Lib 		Library		`xml:"books"`		
}

type BookListXml struct {
	XMLName		xml.Name	`xml:"GoodreadsResponse"`
	Auth		Author		`xml:"author"`
}

func (b Book) String() string {
	return fmt.Sprintf(" Id : %d\n\tTitle with serie : %s\n\tTitle without serie : %s\n",
		b.Id, b.Title, b.NakedTitle)
}

// Main function of the BookXmlParser that fetches the book list XML using an auth id
// and the author's key and parses the formatted XML as well as handling errors

func BookListXmlParser(url string) {

	if XMLdata, err := GetXml(url); err != nil {
		log.Printf("Failed to retrieve XML: %v", err)
	} else {

		var blx BookListXml
		xml.Unmarshal(XMLdata, &blx)

		fmt.Println(blx.Auth.Lib.Books)
	}
}