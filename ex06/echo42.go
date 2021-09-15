package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	e := flag.Bool("e", false, "omit trailing newline")
	s := flag.String("s", " ", "string separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*e {
		fmt.Println()
	}
}
