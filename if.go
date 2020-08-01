package main

import "fmt"

func main() {
	age := 10

	if age > 18 {
		fmt.Println("Senior")
		return
	}

	fmt.Println("Teen")
}
