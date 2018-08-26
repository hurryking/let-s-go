package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	//ch := make(chan int)
	//t := tree.New(1)
	//go Walk(t, ch)
	/*
		for v := range ch {
			fmt.Println(v)
		}
	*/
	fmt.Print(Same(tree.New(1), tree.New(1)))
	fmt.Print(Same(tree.New(2), tree.New(2)))

}
