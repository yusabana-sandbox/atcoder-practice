package main

import "fmt"

func main() {
	a()
}

func a()  {
	var n int
	fmt.Scan(&n)

	// 1桁目取得
	x := n % 10

	switch x {
	case 2,4,5,7,9:
		fmt.Println("hon")
	case 0, 1, 6, 8:
		fmt.Println("pon")
	case 3:
		fmt.Println("bon")
	}
}
