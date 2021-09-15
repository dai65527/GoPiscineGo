package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	dan := 1
	for n >= dan {
		for i := 0; i < dan; i++ {
			fmt.Print("*")
		}
		fmt.Println()
		n -= dan
		dan++
	}
}
