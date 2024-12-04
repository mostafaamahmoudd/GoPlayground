package main

import "fmt"

func loop() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for i <= 3 {
		fmt.Println("Mustafa")
		i++
	}

	j := 0
	for {
		fmt.Print(j)

		if j == 4 {
			break
		}

		j++
	}
}
