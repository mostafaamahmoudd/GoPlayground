package main

import "fmt"

func conditions() {
	if i := 8; i <= 8 {
		fmt.Println("greater than or equal")
	}

	if "Mustafa" == "Nour" || 1 == 1 {
		fmt.Println("Are equal")
	}

	x := 9
	if x >= 0 {
		fmt.Println(x)
	} else {
		fmt.Println(false)
	}
}
