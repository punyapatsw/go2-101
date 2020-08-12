package main

import "fmt"

func main() {
	// primes := []int{2, 3, 5, 7, 11, 13}
	// s := primes[0:4]
	// fmt.Printf("%#v\n", primes[5])
	// fmt.Printf("%#v\n", s)

	// names := []string{"John", "Paul", "George", "Ringo"}
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// s := []int{2, 3, 5, 7, 11, 13}
	// fmt.Printf("length=%d cap=%d\n", len(s), cap(s))

	// s = s[1:4]
	// s = append(s, s...)
	// fmt.Printf("length=%d cap=%d\n", len(s), cap(s))
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Printf("length=%d cap=%d\n", len(s), cap(s))

	// s = s[1:]
	// fmt.Printf("length=%d cap=%d\n", len(s), cap(s))

	// s := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(s[2])
	// s[2] = 123
	// fmt.Println(s[2])

	// s = append(s, 4, 5, 6)
	// fmt.Println(s)

	// data := []int{1, 2, 3, 4}

	// result := append(s, data...)
	// fmt.Println(result)

	// s := []int{}
	// var s []int
	// s := make([]int, 5, 10)
	// fmt.Println(s)

	// if s == nil {
	// 	fmt.Println(nil)
	// } else {
	// 	fmt.Println(s)
	// }

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// for i, v := range pow {
	// 	fmt.Printf("2**%d=%d\n", i, v)
	// }

	for i := range pow {
		fmt.Printf("2**%d\n", i)
	}
}
