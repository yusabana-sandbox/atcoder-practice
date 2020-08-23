package main

import (
	"fmt"
	"strings"
)

func RunA() {
	var s string
	fmt.Scanf("%s", &s)

	//fmt.Print(strings.Count(s, "R"))

	if strings.Contains(s, "RRR") {
		fmt.Println("3")
	} else if strings.Contains(s, "RR") {
		fmt.Println("2")
	} else if strings.Contains(s, "R") {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

func RunAAnswer() {
	var s1 string
	var s2 string
	var s3 string

	fmt.Scanf("%s%s%s", &s1, &s2, &s3)

	b1 := s1 == "R"
	b2 := s2 == "R"
	b3 := s3 == "R"

	if b1 && b2 && b3 {
		fmt.Println("3")
	} else if (b1 && b2) || (b2 && b3) {
		fmt.Println("2")
	} else if b1 || b2 || b3 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
