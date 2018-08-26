package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) float64 {
    z := x / 2
    for i := 0; i < 10; i++ {
        z -= ( z*z - x ) / ( 2*z )
        fmt.Println(z)
    }

    return z
}

func main() {
    n := 3.0
    fmt.Println(sqrt(n))
    fmt.Println(math.Sqrt(n))
}
