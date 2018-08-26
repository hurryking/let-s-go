package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    now := time.Now()

    rand.Seed(now.Unix())
    fmt.Println("My favorite number is", rand.Intn(10), "Now is ", now)
}
