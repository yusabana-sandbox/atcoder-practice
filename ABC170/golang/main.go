package main

import (
	"fmt"
)

func main() {
	a()
}

func a() {
	// makeでスライスを作ってpositionを求めるクロージャもスライスを受け取れるようにする
	// varでlistを定義するだけだと長さが指定できないのでScanで &list[0] のようなアクセスでエラーが出る
	//var list []int
	list := make([]int, 5, 10)

	// スペース区切りなのでそのままScanで取れる
	fmt.Scan(&list[0], &list[1], &list[2], &list[3], &list[4])

	// 0を代入したポジションを返す
	// https://gawawa124.hatenablog.com/entry/2015/04/08/193237
	pos := func(arr []int, e int) int {
		for i, v := range arr {
			if v == e { return i + 1 }
		}
		return -1
	}(list, 0)

	fmt.Println(pos)
}
