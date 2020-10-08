package main

import (
	"fmt"
	"sort"
)

// golang 標準入力から空白区切りの数列を読み込む | ikapblog
// https://blog.ikappio.com/golang-read-space-separated-integers-from-stdin/
func RunB() {
	var n int
	fmt.Scanf("%d", &n)

	// sliceを初期化してscanfで順次読み込む
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &l[i])
	}
	// 昇順で並び替える
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })

	result := 0

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if l[i] != l[j] && l[i] != l[k] && l[j] != l[k] &&
					l[i]+l[j] > l[k] {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}
