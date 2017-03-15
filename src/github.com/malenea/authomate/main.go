package main

import (
		"os"
		"strings"
		"bytes"
		"fmt"
		"path/filepath"
		"strconv"
)

func BookListConcat(id string) string {
	var BookList bytes.Buffer

	BookList.WriteString("https://www.goodreads.com/author/list.xml?key=")
	BookList.WriteString("kDkKnUxiz8cRBJhVjrtSA")
	BookList.WriteString("&id=")
	BookList.WriteString(id)

	return BookList.String()
}

func main() {
	if len(os.Args) > 1 {

		var adress, id string

		adress = os.Args[1]

		_, err := strconv.ParseInt(adress, 10, 0)
		if err == nil {
			id =  adress
		} else if strings.Contains(adress, "http://") || strings.Contains(adress, "https://") {
			id =  AuthorXmlParserFromUrl(adress)
		} else if strings.Compare(filepath.Ext(adress), ".xml") == 0 {
			id = AuthorXmlParserFromFile(adress) // Clean id format "id id id id id"
		}

		BookListQuery := BookListConcat(id)
		BookListXmlParser(BookListQuery)

		return
	}
	fmt.Println("Authomate takes as parameter either the XML answer of a book review")
	fmt.Println("or the .xml file of the book review")
}