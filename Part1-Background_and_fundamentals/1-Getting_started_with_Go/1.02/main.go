package main

import (
	"fmt"
)

func getStrings() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := getStrings()
	fmt.Println(n1, n2)
}
