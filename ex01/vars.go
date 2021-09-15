package main

import "fmt"

type FortyTwo struct{}

func main() {
	var c chan int

	vars := []interface{}{
		"42",
		uint(42),
		int(42),
		uint8(42),
		int16(42),
		uint32(42),
		int64(42),
		false,
		float32(42),
		float64(42),
		complex64(42),
		complex(42, 21),
		FortyTwo{},
		[1]int{42},
		map[string]int{"42": 42},
		(*int)(nil),
		[]int{},
		c,
		nil,
	}

	for _, v := range vars {
		putvar(v)
	}
}

func putvar(v interface{}) {
	typeName := fmt.Sprintf("%T", v)
	prepos := "a"

	if typeName == "*int" {
		fmt.Printf("%p is an %s.\n", v, typeName)
		return
	}

	if typeName[0] == 'i' {
		prepos = "an"
	}
	fmt.Printf("%v is %s %s.\n", v, prepos, typeName)
}
