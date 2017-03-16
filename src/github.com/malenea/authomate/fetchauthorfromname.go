package main

import (
		"io/ioutil"
		"net/http"
		"fmt"
		"strings"
		"bytes"
		"regexp"
		"log"
)

// Concat the URL for the author's id request using the author's name

func FetchAuthorConcat(name string) string {
	var AuthorName bytes.Buffer

	name = strings.TrimSpace(name)
	name = strings.Replace(name, " ", "_", -1)
	AuthorName.WriteString("https://www.goodreads.com/api/author_url/")
	AuthorName.WriteString(name)
	AuthorName.WriteString("?key=kDkKnUxiz8cRBJhVjrtSA")

	return AuthorName.String()
}

// Function that allows to get a goodreaders' author's id in a XML format author's review

func GetIdFromXml(data, fdelim, bdelim string, oc int) string {
	datatab := strings.Split(data, fdelim)
	if len(datatab) <= oc {
		return ""
	}

	tmpdata := datatab[oc]
	datatab = strings.Split(tmpdata, bdelim)
	if len(datatab) == 1 {
		return ""
	}

	return datatab[0]
}

// Main function that will fetch the author's id from a XML format and return it as
// a string

func FetchAuthorFromName(name string) string {
	strings.Join(strings.Fields(name), " ")
	inwhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	name = inwhitespace.ReplaceAllString(name, " ")
	concaturl := FetchAuthorConcat(name)

	xmlresp, err := http.Get(concaturl)
	if err != nil {
		log.Printf("Failed to http.Get XML: %v", err)
	} else {
		defer xmlresp.Body.Close()
		xmldata, err := ioutil.ReadAll(xmlresp.Body)
		if err != nil {
			log.Printf("Failed to read XML: %v", err)
		}

		return GetIdFromXml(string(xmldata), "author id=\"", "\"", 1)
	}

	return ""
}