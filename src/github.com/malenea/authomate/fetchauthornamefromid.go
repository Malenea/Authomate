package main

import (
		"io/ioutil"
		"net/http"
		"strings"
		"bytes"
		"regexp"
		"log"
)

// Concat the URL for the author's id request using the author's name

func FetchAuthorIdConcat(key, id string) string {
	var AuthorId bytes.Buffer

	id = strings.TrimSpace(id)
	AuthorId.WriteString("https://www.goodreads.com/author/show.xml?key=")
	AuthorId.WriteString(key)
	AuthorId.WriteString("&id=")
	AuthorId.WriteString(id)

	return AuthorId.String()
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

func FetchAuthorNameFromId(key, id string) string {
	strings.Join(strings.Fields(id), " ")
	inwhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	id = inwhitespace.ReplaceAllString(id, " ")
	concaturl := FetchAuthorIdConcat(key, id)

	xmlresp, err := http.Get(concaturl)
	if err != nil {
		log.Printf("Failed to http.Get XML: %v", err)
	} else {
		defer xmlresp.Body.Close()
		xmldata, err := ioutil.ReadAll(xmlresp.Body)
		if err != nil {
			log.Printf("Failed to read XML: %v", err)
		}

		return GetStrFromXml(string(xmldata), "<name>", "</name>", 1)
	}

	return ""
}