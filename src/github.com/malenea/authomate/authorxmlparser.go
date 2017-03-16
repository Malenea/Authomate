package main

import (
		"os"
		"io/ioutil"
		"encoding/xml"
		"fmt"
		"log"
)

// Structures for XML beacons

type Writer struct {
	XMLName		xml.Name	`xml:"author"`
	Id 			int			`xml:"id"`
	Name		string 		`xml:"name"`
}

type AuthPool struct {
	XMLName		xml.Name	`xml:"authors"`
	Auth 		[]Writer	`xml:"author"`
}

type SingleBook struct {
	XMLName		xml. Name 	`xml:"book"`
	Pool 		AuthPool	`xml:"authors"`
}

type AuthorXml struct {
	XMLName		xml.Name	`xml:"GoodreadsResponse"`
	Work		SingleBook	`xml:"book"`
}

func (w Writer) String() string {
	return fmt.Sprintf("%d", w.Id)
}

// Functions of the AuthorXmlParser that creates the author key using an auth id
// from a book's review XML page or file

func AuthorXmlParserFromUrl(url string) string {

	if XMLdata, err := GetXml(url); err != nil {
		log.Printf("Failed to retrieve XML: %v", err)
	} else {
		var ax AuthorXml
		xml.Unmarshal(XMLdata, &ax)
		return fmt.Sprintf("%s", ax.Work.Pool.Auth)
	}

	return ""
}

func AuthorXmlParserFromFile(path string) string {
	XMLfile, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open XML file: %v", err)
		return ""
	}
	defer XMLfile.Close()

	XMLdata, _ := ioutil.ReadAll(XMLfile)

	var ax AuthorXml
	xml.Unmarshal(XMLdata, &ax)

	return fmt.Sprintf("%s", ax.Work.Pool.Auth)
}