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

	if n == 'a' { }


	msg := "world"
	fmt.Printf("hello, %s!", msg)

	fmt.Println(n + math.Pow(n, 2)+ math.Pow(n, 3))
}
