package main

import (
	"fmt"
)

func main() {
	a()
}

func a() {
	var list [5]int

	fmt.Scan(&list[0], &list[1], &list[2], &list[3], &list[4]) // スペース区切りなのでそのままScanで取れる

	// 0を代入したポジションを返す
	pos := func(arr [5]int, e int) int {
		for i, v := range arr {
			if v == e { return i + 1 }
		}
		return -1
	}(list, 0)

	fmt.Println(pos)
}
