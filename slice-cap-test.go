package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	a := make([]int, 1, 2)
	printSlice(a)

	a = append(a, 2)
	printSlice(a)

	a = append(a, 32)
	printSlice(a)
}
