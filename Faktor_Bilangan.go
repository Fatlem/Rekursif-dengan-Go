package main

import (
	"fmt"
	"strconv"
)

func printFactors(n int, current int) {
	if current > n {
		return
	}
	if n%current == 0 {
		fmt.Printf("%d ", current)
	}
	printFactors(n, current+1)
}

func main() {
	var input string
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&input)

	N, err := strconv.Atoi(input)
	if err != nil || N <= 0 {
		fmt.Println("Masukan bilangan bulat positif yang valid.")
		return
	}

	fmt.Printf("Faktor dari %d adalah: ", N)
	printFactors(N, 1)
	fmt.Println()
}
