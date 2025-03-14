package main

import (
	"fmt"
)

func power(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var x, y int

	fmt.Print("Masukan nilai (x): ")
	fmt.Scanln(&x)
	fmt.Print("Masukan nilai (y): ")
	fmt.Scanln(&y)

	hasil := power(x, y)
	fmt.Printf("%d dipangkatkan %d adalah: %d\n", x, y, hasil)
}
