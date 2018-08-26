package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {12.41214, -74.25433},
	"Google":    {12.41214, -74.25433},
}

var n = map[string]uint8{
	"A": 12,
	"B": 23,
}

func main() {
	fmt.Println(m)
	n["C"] = 45
	fmt.Println(n)
}
