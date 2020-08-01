package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(hypot(1, 5))
	fmt.Println(compute(hypot, 7))
	fmt.Println(compute(math.Pow, 7))
}

func compute(fn func(float64, float64) float64, x int) float64 {
	return fn(float64(x), 10.1) + float64(x)
}
