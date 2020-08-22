package main

import "fmt"

type OnePlus struct {
	Version string
}

func main() {
	var i interface{} = "hello"
	i = &OnePlus{Version: "7 Pro"}
	switch v := i.(type) {
	case int:
		fmt.Printf("%T %d\n", v, v)
	case string:
		fmt.Printf("%T %s\n", v, v)
	case OnePlus, *OnePlus:
		fmt.Printf("%T %v yay my new phone\n", v, v)
	default:
		fmt.Println("undefined type")
	}
}
