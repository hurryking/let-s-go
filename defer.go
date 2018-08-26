package main

import "fmt"

func hello() {
    fmt.Print("a")
}

func main() {
    hello()
    defer fmt.Print("b") 
    fmt.Print("c")
}
