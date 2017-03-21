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

func FetchAuthorConcat(key, name string) string {
	var AuthorName bytes.Buffer

	name = strings.TrimSpace(name)
	name = strings.Replace(name, " ", "_", -1)
	AuthorName.WriteString("https://www.goodreads.com/api/author_url/")
	AuthorName.WriteString(name)
	AuthorName.WriteString("?key=")
	AuthorName.WriteString(key)

	return AuthorName.String()
}

// Function that allows to get a special string in a XML format

func GetStrFromXml(data, fdelim, bdelim string, oc int) string {
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

func FetchAuthorFromAuthorUrl(url string) (string, string) {
	xmlresp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to http.Get XML: %v", err)
	} else {
		defer xmlresp.Body.Close()
		xmldata, err := ioutil.ReadAll(xmlresp.Body)
		if err != nil {
			log.Printf("Failed to read XML: %v", err)
		}

		id := GetStrFromXml(string(xmldata), "<author id=\"", "\">", 1)
		name := GetStrFromXml(string(xmldata), "<name><![CDATA[", "]]></name>", 1)
		return id, name
	}

	return "", ""
}

func FetchAuthorFromName(key, name string) string {
	strings.Join(strings.Fields(name), " ")
	inwhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	name = inwhitespace.ReplaceAllString(name, " ")
	concaturl := FetchAuthorConcat(key, name)

	xmlresp, err := http.Get(concaturl)
	if err != nil {
		log.Printf("Failed to http.Get XML: %v", err)
	} else {
		defer xmlresp.Body.Close()
		xmldata, err := ioutil.ReadAll(xmlresp.Body)
		if err != nil {
			log.Printf("Failed to read XML: %v", err)
		}

		return GetStrFromXml(string(xmldata), "author id=\"", "\"", 1)
	}

	return ""
}