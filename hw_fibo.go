package main

import "fmt"

func fibo() func() int {
	pos0 := 0
	pos1 := 0
	return func() int {
		if pos0 == 0 && pos1 == 0 {
			defer func() {
				pos1 = 1
			}()
			// return pos0
		}
		tmp := pos0
		pos0 = pos1
		pos1 += tmp
		return pos0
	}
}

func main() {
	f := fibo()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
