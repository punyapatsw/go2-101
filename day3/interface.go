package main

import "fmt"

type square struct {
	width float64
}

func (s square) Area() float64 {
	return s.width * s.width
}

type rec struct {
	width float64
	high  float64
}

func (r rec) Area() float64 {
	return r.width * r.high
}

type Areaer interface {
	Area() float64
}

func showArea(s Areaer) {
	fmt.Println(s.Area())
}

func main() {
	a := square{width: 4}
	showArea(a)
	b := rec{width: 4, high: 3}
	showArea(b)
}
