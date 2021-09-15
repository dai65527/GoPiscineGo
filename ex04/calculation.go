package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG = "Arguments is invalid."

func main() {
	s, ok := calulationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}

func calulationStr(args []string) (string, bool) {
	if len(args) != 3 {
		return "", false
	}

	a, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return "", false
	}

	b, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil || b == 0 {
		return "", false
	}

	s := fmt.Sprintf("sum: %d\n", a+b)
	s += fmt.Sprintf("difference: %d\n", a-b)
	s += fmt.Sprintf("product: %d\n", a*b)
	s += fmt.Sprintf("quotient: %d\n", a/b)
	return s, true
}
