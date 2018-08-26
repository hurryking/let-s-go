package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    i := 42
    f64 := i
    u := uint(f64)
    fmt.Println(x, y, z, i, f64, u)
}
