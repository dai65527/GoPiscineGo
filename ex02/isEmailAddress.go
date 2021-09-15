package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
	}

	r := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)

	ok := true
	for _, address := range os.Args[1:] {
		if len(address) > 256 || !r.Match([]byte(address)) {
			fmt.Printf("%s is NOT a valid email address.\n", address)
		} else {
			fmt.Printf("%s is a valid email address.\n", address)
		}
	}

	if !ok {
		fmt.Println("")
	}
}
