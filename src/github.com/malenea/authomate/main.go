package main

import (
		"os"
		"strings"
		"bytes"
		"fmt"
		"path/filepath"
		"strconv"
)

// Concat the URL for the author's book list request using the author's id

func BookListConcat(id string) string {
	var BookList bytes.Buffer

	BookList.WriteString("https://www.goodreads.com/author/list.xml?key=")
	BookList.WriteString("kDkKnUxiz8cRBJhVjrtSA")
	BookList.WriteString("&id=")
	BookList.WriteString(id)

	return BookList.String()
}

// GetId function that allows input parameter, direct int id / URL XML / .xml

func GetId(adress string) string {
	var id string
	_, err := strconv.ParseInt(adress, 10, 0)
	if err == nil {
		id =  adress
	} else if strings.Contains(adress, "http://") || strings.Contains(adress, "https://") {
		id =  AuthorXmlParserFromUrl(adress)
	} else if strings.Compare(filepath.Ext(adress), ".xml") == 0 {
		id = AuthorXmlParserFromFile(adress)
	}
	id = strings.Replace(id, "[", "", -1)
	id = strings.Replace(id, "]", "", -1)
	return id
}

// Main function

func main() {
	if len(os.Args) > 1 {

		var adress, id string

		for it := range os.Args {
		adress = os.Args[it]

		id = GetId(adress)
		idArray := strings.Fields(id)

		for _, each := range idArray {
			BookListQuery := BookListConcat(each)
			BookListXmlParser(BookListQuery)
		}
	}

		return
	}
	fmt.Println("Authomate takes as parameter either the XML answer of a book review")
	fmt.Println("or the .xml file of the book review")
}