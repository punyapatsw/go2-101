package main

import "fmt"

func main() {
	var i interface{} = "2"
	s := i.(string)
	// fmt.Println(s)
	if s, ok := i.(string); ok {
		fmt.Println(s, ok)
	}

	x, _ := i.(string)
	fmt.Println(x)
}
