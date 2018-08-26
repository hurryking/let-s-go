package main

import "fmt"

func main() {
	ch := make(chan []int, 2)

	ch <- []int{1, 2, 3}
	ch <- []int{3, 4, 5}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
