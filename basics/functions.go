package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func multipleVals() (int, int) {
	return 1, 2
}

func functions() {
	fmt.Println(sum(1, 2))

	fmt.Println(plusplus(1, 2, 3))

	a, b := multipleVals()
	fmt.Println(a, b)
}
