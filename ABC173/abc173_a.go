package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n < 1000 {
		fmt.Println(1000 - n)
	} else {
		for n > 1000 {
			n = n - 1000
		}
		fmt.Println(1000 - n)
	}
}
