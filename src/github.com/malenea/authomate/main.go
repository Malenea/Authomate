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

func BookListConcat(key, id string) string {
	var BookList bytes.Buffer

	BookList.WriteString("https://www.goodreads.com/author/list.xml?key=")
	BookList.WriteString(key)
	BookList.WriteString("&id=")
	BookList.WriteString(id)

	return BookList.String()
}

// GetId function that allows input parameter, direct int id / URL XML / .xml

func GetId(key, value string) []string {
	var idarray []string

	_, err := strconv.ParseInt(value, 10, 0)
	if err == nil {
		idarray = append(idarray, value)
		return idarray
	} else if strings.Contains(value, "http://") || strings.Contains(value, "https://") {
		idarray =  AuthorXmlParserFromUrl(value)
	} else if strings.Compare(filepath.Ext(value), ".xml") == 0 {
		idarray = AuthorXmlParserFromFile(value)
	} else {
		id := FetchAuthorFromName(key, value)
		idarray = append(idarray, id)
	}

	return idarray
}

// Main function

func main() {
	key := "0"
	if len(os.Args) > 1 {

		var adress string
		args := os.Args[1:]

		for it := range args {
			adress = args[it]

			idarray := GetId(key, adress)

			for _, eachid := range idarray {
				booklistquery := BookListConcat(key, eachid)
				booklist := BookListXmlParser(booklistquery)

				for _, eachbook := range booklist {
					fmt.Println(eachbook)
				}
			}
		}
	return
	}

	fmt.Println("Authomate takes as parameter either the XML answer of a book review")
	fmt.Println("or the .xml file of the book review")
}