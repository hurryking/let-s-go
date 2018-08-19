package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }

    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
        cmplx.Sqrt(3+4i),
    )
}
