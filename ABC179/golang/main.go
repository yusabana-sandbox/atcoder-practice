package main

import "fmt"

func main() {
	a()
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
