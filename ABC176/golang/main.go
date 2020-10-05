package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//a()
	b()
}

func b() {
	var input string
	var sum int = 0
	var nums []string

	fmt.Scan(&input)

	nums = strings.Split(input, "")

	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		sum += num
	}

	if sum%9 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func a() {
	var n float64
	var x float64
	var t float64

	fmt.Scanf("%f %f %f", &n, &x, &t)
	num := math.Ceil(n / x)
	fmt.Printf("%.0f", t*num)
}
