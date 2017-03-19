package main

import (
		"os"
		"testing"
		"fmt"
)

// Retrieving dev key from environment variable

var testkey = os.Getenv("GR_DEVKEY")

// Testing fetching functions

func TestFetchAuthorFromName(t *testing.T) {
	var id string
	var JKRname, SKname string = "J.K. Rowling", "Stephen King"
	var JKRid, SKid string = "1077326", "3389"

	id = FetchAuthorFromName(testkey, JKRname)
	if id == JKRid {
		fmt.Println("Got: ", id, " for ", JKRname)
	} else {
		t.Error("Expected ", JKRid, ", got ", id)
	}
	id = FetchAuthorFromName(testkey, SKname)
	if id == SKid {
		fmt.Println("Got: ", id, " for ", SKname)
	} else {
		t.Error("Expected ", SKid, ", got ", id)
	}
}

func TestFetchAuthorFromId(t *testing.T) {
	var name string
	var JKRname, SKname string = "J.K. Rowling", "Stephen King"
	var JKRid, SKid string = "1077326", "3389"

	name = FetchAuthorNameFromId(testkey, JKRid)
	if name == JKRname {
		fmt.Println("Got: ", name, " for ", JKRid)
	} else {
		t.Error("Expected ", JKRname, ", got ", name)
	}
	name = FetchAuthorNameFromId(testkey, SKid)
	if name == SKname {
		fmt.Println("Got: ", name, " for ", SKid)
	} else {
		t.Error("Expected ", SKname, ", got ", name)
	}
}

// Testing sorting function

func TestGetId(t *testing.T) {
	var namearray, idarray []string
	var JKRid, SKid string = "1077326", "3389"
	var path string = os.Getenv("PWD") + "/../../../../resources/Goodreads.xml"
	var url string = FetchAuthorConcat(testkey, "Stephen_King")

	idarray, namearray = GetId(testkey, JKRid)
	fmt.Println(namearray, ": ", idarray, " from ", JKRid)
	idarray, namearray = GetId(testkey, SKid)
	fmt.Println(namearray, ": ", idarray, " from ", SKid)
	idarray, namearray = GetId(testkey, path)
	fmt.Println(namearray, ": ", idarray, " from ", path)
	idarray, namearray = GetId(testkey, url)
	fmt.Println(namearray, ": ", idarray, " from ", url)
}

// Testing concat function

func TestBookListConcat(t *testing.T) {
	var JKRid, SKid string = "1077326", "3389"
	testids := [2]string{JKRid, SKid}

	for _, eachid := range testids {
		concaturl := BookListConcat(testkey, eachid, "1")
		fmt.Println(concaturl)
	}
}

func TestIteratePages(t *testing.T) {
	idarray := []string{"Stephen King", "J.K. Rowling"}
	namearray := []string{"3389", "1077326"}

	IteratePages(idarray, namearray, testkey)
}