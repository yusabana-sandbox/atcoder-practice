package main

import (
	"fmt"
	"math"
)

func main() {
	a()
}

func a() {
	var n float64
	fmt.Scan(&n)
	fmt.Println(n + math.Pow(n, 2)+ math.Pow(n, 3))
}
