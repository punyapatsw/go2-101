package main

import "fmt"

const Pi = 3.14

const (
	A = 1
	_ = iota * 2
	Monday
	Tuesday
	Wednesday
	Test = "hello"
	Thursday
)

func main() {
	const world = "world"
	fmt.Printf("Hello %T\n", Pi)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)
	fmt.Println(A, Monday, Tuesday, Wednesday, Test, Thursday)
}
