package main

import (
	"fmt"
	"strconv"
)

func printOddNumbers(n int, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", current)
	printOddNumbers(n, current+2)
}

func main() {
	var input string
	fmt.Print("Masukan bilangan bulat positif N: ")
	fmt.Scanln(&input)

	N, err := strconv.Atoi(input)
	if err != nil || N <= 0 {
		fmt.Println("Masukan bilangan bulat positif yang valid.")
		return
	}

	fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah: ", N)
	printOddNumbers(N, 1)
	fmt.Println()
}
