package main

import (
		"fmt"
		"os"
		"io/ioutil"
		"encoding/xml"
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

// Function that returns only the author's names of each author from the []Writer
// in a []string

func WriterStructToArray(writerarray []Writer) ([]string, []string) {
	var resid, resname []string

	for _, each := range writerarray {
		resid = append(resid, fmt.Sprintf("%d ", each.Id))
		resname = append(resname, fmt.Sprintf("%s ", each.Id))
	}
	return resid, resname
}

// Functions of the AuthorXmlParser that creates the author key using an auth id
// from a book's review XML page or file

func AuthorXmlParserFromUrl(url string) ([]string, []string) {

	if XMLdata, err := GetXml(url); err != nil {
		log.Printf("Failed to retrieve XML: %v", err)
	} else {
		var ax AuthorXml
		xml.Unmarshal(XMLdata, &ax)

		return WriterStructToArray(ax.Work.Pool.Auth)
	}

	return nil, nil
}

func AuthorXmlParserFromFile(path string) ([]string, []string) {
	XMLfile, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open XML file: %v", err)
		return nil, nil
	}
	defer XMLfile.Close()

	XMLdata, _ := ioutil.ReadAll(XMLfile)

	var ax AuthorXml
	xml.Unmarshal(XMLdata, &ax)

	return WriterStructToArray(ax.Work.Pool.Auth)
}