package main

import (
	"fmt"
)

func main() {
	//RunA()
	RunB()
}

func RunB() {
	var N, D int
	fmt.Scanf("%d %d", &N, &D)

	cnt := 0
	//a := make([][]int, N)

	//for i, _ := range a {
	//	var x, y int
	//	fmt.Scan(&x, &y)
	//	a[i] = []int{x, y}
	//}
	//
	//for _, it := range a {
	//	if it[0]*it[0]+it[1]*it[1] <= D*D {
	//		cnt++
	//	}
	//}

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		if x*x+y*y <= D*D {
			cnt++
		}
	}

	fmt.Println(cnt)
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
