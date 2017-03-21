package main

import (
		"net/http"
		"io/ioutil"
		"fmt"
		"log"
)

// Will allow us to make the http request, taking as parameter the request url
// and return formatted XML and error code

func GetXml(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 404 {
			log.Printf("Author could not be found, check author's name again")
		}
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Html body error: %v", err)
	}

	return data, nil
}