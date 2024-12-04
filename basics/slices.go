package main

import (
	"fmt"
	"slices"
)

func slices() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	slices.Reverse(numbers)
	fmt.Println(numbers)

	slices.Sort(numbers)
	fmt.Println(numbers)

	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println(s, s == nil, cap(s))

	s[0] = "Mustafa"
	s[1] = "Mahmoud"
	s[2] = "Msahel"

	s = append(s, "Mustafa", "Mahmoud")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[1:3]
	fmt.Println(l)

	l = s[:3]
	fmt.Println(l)

	l = s[1:]
	fmt.Println(l)
}
