package main

import (
	"fmt"
	"sort"
)

// golang 標準入力から空白区切りの数列を読み込む | ikapblog
// https://blog.ikappio.com/golang-read-space-separated-integers-from-stdin/
// この回答例参考になる
//   https://atcoder.jp/contests/abc175/submissions/16563798
func RunB() {
	var n int
	fmt.Scanf("%d", &n)

	// sliceを初期化してscanfで順次読み込む
	l := make([]int, n)
	for i := 0; i < n; i++ { // このfor文でもいける => for i := range l {
		fmt.Scanf("%d", &l[i])
	}
	// 昇順で並び替える
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	// 以下のいずれかでも良さそう
	//sort.Ints(l)
	//sort.Sort(sort.IntSlice(l))
	//sort.IntSlice(l).Sort()

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
