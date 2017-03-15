package main

import (
		"os"
		"fmt"
)

func main() {
	if len(os.Args) > 1 {

		XmlParser(os.Args[1], "kDkKnUxiz8cRBJhVjrtSA")

		return
	}
	fmt.Println("Please provide an argument")
}