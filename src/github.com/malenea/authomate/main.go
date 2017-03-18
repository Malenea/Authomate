package main

import (
		"strings"
		"bytes"
		"path/filepath"
		"strconv"
		"flag"
		"log"
)

// Concat the URL for the author's book list request using the author's id

func BookListConcat(key, id string, page string) string {
	var BookList bytes.Buffer

	BookList.WriteString("https://www.goodreads.com/author/list.xml?key=")
	BookList.WriteString(key)
	BookList.WriteString("&id=")
	BookList.WriteString(id)
	BookList.WriteString("&page=")
	BookList.WriteString(page)

	return BookList.String()
}

// GetId function that allows input parameter, direct int id / URL XML / .xml

func SortUrlType(url string) bool {
	if strings.Contains(url, "/author_url/") {
		return true
	} else {
		return false
	}
}

func GetId(key, value string) ([]string, []string) {
	var idarray, namearray []string
	var id, name string

	_, err := strconv.ParseInt(value, 10, 0)
	if err == nil {
		idarray = append(idarray, value)
		namearray = append(namearray, FetchAuthorNameFromId(key, value))
		return idarray, namearray
	} else if strings.Contains(value, "http://") || strings.Contains(value, "https://") {
		if SortUrlType(value) {
			id, name = FetchAuthorFromAuthorUrl(value)
			idarray = append(idarray, id)
			namearray = append(namearray, name)
		} else {
			idarray, namearray =  AuthorXmlParserFromUrl(value)
		}
	} else if strings.Compare(filepath.Ext(value), ".xml") == 0 {
		idarray, namearray = AuthorXmlParserFromFile(value)
	} else {
		id = FetchAuthorFromName(key, value)
		idarray = append(idarray, id)
		namearray = append(namearray, value)
	}

	return idarray, namearray
}

// Function to reformat map[string]interface{}

func ToMstringInt(name string, books []string) map[string]interface{} {
	var result map[string]interface{}

	if result == nil {
		result = make(map[string]interface{})
	}

	_, ok := result[name]
	if !ok {
		bookvalues := make ([]interface{}, len(books))
		for i, b := range books {
			bookvalues[i] = b
		}
		result[name] = bookvalues
	}

	return result
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

			idarray, namearray := GetId(key.Value, adress)
			IteratePages(idarray, namearray, key.Value)
		}
	return
	}

	log.Printf("Usage of Authomate : authomate -key={your_key} Args[...]")
	log.Printf("Where Args are the author's ids or names, or the xml files or urls of")
	log.Printf("one or more of the author's book review")
}