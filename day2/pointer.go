package main

import "fmt"

func main() {
	// i := 42
	// fmt.Println("Address :", &i)

	// var p *int
	// p = &i
	// fmt.Println("Value :", p)
	// fmt.Println("Value inside :", *p)

	// i = 41
	// fmt.Println("Value :", p)
	// fmt.Println("Value inside :", *p)

	// *p = 45
	// fmt.Println("Value :", i)
	// fmt.Println("Value inside :", *p)

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	// p = 10
	// fmt.Println(j)

}
