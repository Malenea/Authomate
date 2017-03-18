package main

import (
		"fmt"
		"log"
		"strconv"
		"sync"
)

var wg sync.WaitGroup

// Function to iterate through pages and print the result

func PrintPages(key, eachid string, pageit int) {
	defer wg.Done()
	booklistquery := BookListConcat(key, eachid, strconv.Itoa(pageit))
	booklist, _ := BookListXmlParser(booklistquery)

	// Commented code is formated map[string]interface{}

	/*result := ToMstringInt(namearray[i], booklist)
	fmt.Println(result)*/

	// This part is just a nice output of the results (readable)

	for _, eachbook := range booklist {
		fmt.Println(eachbook)
	}
}

func IteratePages(idarray, namearray []string, key string) {
	for i, eachid := range idarray {

		fmt.Println(namearray[i])
		page := "1"

		booklistquery := BookListConcat(key, eachid, page)
		_, total := BookListXmlParser(booklistquery)

		tmp, err := strconv.Atoi(total)
		if err != nil {
			log.Printf("Error occured from page count: %v", err)
			break
		}
		sum := 1
	for i := 30; i < tmp; i += 30 {
			sum +=  1
		}

		wg.Add(sum)

		for pageit := 1; pageit <= sum; pageit++ {
			go PrintPages(key, eachid, pageit)
		}

		wg.Wait()

		fmt.Println("")
	}
}