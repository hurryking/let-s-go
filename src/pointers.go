package main

import "fmt"

func main() {
    i, j := 12, 14

    p := &i
    fmt.Println(*p)
    *p = 21
    fmt.Println(i)

    p = &j
    fmt.Println(*p)
}
