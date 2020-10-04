package main

import "fmt"

func main() {
	//a()
	b()
}

func b() {
	zCnt := 0
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var d1, d2 int
		fmt.Scanf("%d %d", &d1, &d2)

		if d1 == d2 {
			zCnt++
			if zCnt == 3 {
				fmt.Println("Yes")
				return
			}
		} else {
			zCnt = 0
		}
	}
	fmt.Println("No")
}

func a() {
	// https://www.it-swarm-ja.tech/ja/string/golang文字列の最後のx文字を取得する方法は？/1049363837/
	var str string
	fmt.Scan(&str)

	// indexを指定して取得パターン
	//fmt.Println(string(str[len(str)-1]))
	// 範囲指定を使ってで取得
	last := str[len(str)-1:]

	if last == "s" {
		fmt.Println(str + "es")
	} else {
		fmt.Println(str + "s")
	}
}
