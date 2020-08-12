package main

import (
	"fmt"

	"named"
)

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y := named.Split(1, 2)
	fmt.Println(a, b, x, y)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
