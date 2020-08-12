package main

import "fmt"

func show(s [3]int) {
	fmt.Println(s)
}

func main() {
	var a [3]int
	a[0] = 10
	a[1] = 11

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	show(a)
	show(primes[0:2])
}
