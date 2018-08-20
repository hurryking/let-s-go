package main

import "fmt"

func main() {
    sum := 1
    i := 1
    for sum < 1000 {
        sum += sum
        fmt.Println(i)
        i++
    }

    fmt.Println(sum)
}
