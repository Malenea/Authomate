package main

import (
		"strings"
		"bytes"
		"fmt"
		"path/filepath"
		"strconv"
		"flag"
		"log"
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

// Main function with flags management

type StringFlag struct {
	IsSet	bool
	Value	string
}

func (sf *StringFlag) Set(value string) error {
	sf.Value = value
	sf.IsSet = true
	return nil
}

func (sf *StringFlag) String() string {
	return sf.Value
}

var key StringFlag

func main() {
	flag.Var(&key, "key", "Your dev key")
	flag.Parse()

	if !key.IsSet {
		log.Printf("Please provide a dev key")
		log.Printf("Usage of Authomate : authomate -key={your_key} Args[...]")
		return
	}

	args := flag.Args()

	if len(args) >= 1 {

		var adress string

		for it := range args {
			adress = args[it]

			idarray := GetId(key.Value, adress)

			for _, eachid := range idarray {
				booklistquery := BookListConcat(key.Value, eachid)
				booklist := BookListXmlParser(booklistquery)

				for _, eachbook := range booklist {
					fmt.Println(eachbook)
				}
			}
		}
	return
	}

	log.Printf("Usage of Authomate : authomate -key={your_key} Args[...]")
	log.Printf("Where Args are the author's ids or names, or the xml files or urls of")
	log.Printf("one or more of the author's book review")
}