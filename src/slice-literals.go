package main

import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    q[0] = 12
    fmt.Println(q)

    r := []bool{true, false, true, true, true, false}
    fmt.Println(r)

    s := []struct {
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {4, true},
        {5, true},
        {6, true},
    }
    s[4].i = 7
    fmt.Println(s)
}
