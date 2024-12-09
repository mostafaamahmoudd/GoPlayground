package main

import (
	"fmt"
	"maps"
)

func main() {
	ages := make(map[string]int)
	ages["Mustafa"] = 25
	ages["Mahmoud"] = 26
	fmt.Println(ages)

	mAge := ages["Mahmoud"]
	fmt.Println(mAge)

	fmt.Println(len(ages))

	delete(ages, "Mustafa")
	fmt.Println(ages)

	clear(ages)
	fmt.Println(ages)

	n := map[string]int {
		"foo": 1,
		"bar": 2,
	}

	m := map[string]int {
		"foo": 1,
		"bar": 2,
	}

	if maps.Equal(n, m) {
		fmt.Println("n == m")
	} else {
		fmt.Println("n != m")
	}
}