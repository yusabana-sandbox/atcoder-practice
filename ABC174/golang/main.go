package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//RunA()
	RunB()
}

var sc = bufio.NewScanner(os.Stdin)
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func RunB() {
	var N, D int
	fmt.Scanf("%d %d", &N, &D)

	sc.Split(bufio.ScanWords)

	cnt := 0
	distance := D * D

	for i := 0; i < N; i++ {
		x, y := nextInt(), nextInt()

		if x*x+y*y <= distance {
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
