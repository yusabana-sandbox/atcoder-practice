package main

import (
	"fmt"
	"math"
)

func main() {
	//RunA()
	RunB()
}

func RunB() {
	var count int
	var N, D float64
	fmt.Scanf("%f %f", &N, &D)

	for i := float64(0); i < N; i++ {
		var x, y float64
		fmt.Scanf("%f %f", &x, &y)

		distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

		if distance <= D {
			count++
		}
	}

	fmt.Println(count)
}

func RunA() {
	var n int
	fmt.Scanf("%d", &n)

	if n >= 30 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}



