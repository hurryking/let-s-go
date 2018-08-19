package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe bool = false
    MaxInt uint = 1<<64 - 1
    MaxInt8 uint8 = 1<<8 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
    i = 1
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
    fmt.Printf("Type: %T Value: %#V\n", MaxInt8, MaxInt8)
    fmt.Printf("Type: %T Value: %#V\n", i, i)
}
