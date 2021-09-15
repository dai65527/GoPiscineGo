package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	e := flag.Bool("n", false, "omit trailing newline")
	s := flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*e {
		fmt.Println()
	}
}
