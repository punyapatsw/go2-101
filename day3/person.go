package main

import "fmt"

type Person struct {
	name    string
	friends map[string]int
	test    []int
}

func (p Person) Walk() {
	fmt.Println(p.name, "walking")
}
func (p Person) Eat() {
	fmt.Println(p.name, "eating")
}
func (p Person) Greeting() {
	fmt.Println("Hello", p.name)
}

func (p Person) Name() string {
	return p.name
}

// func (p Person) SetName(s string) Person {
// 	p.name = s
// 	return p
// }

func (p *Person) SetName(s string) {
	p.name = s
	p.friends["B"] = 6
	p.test[0] = 3
}

// func main() {
// 	p := Person{"A"}
// 	p.Walk()
// 	p.Eat()
// 	p.Greeting()
// 	fmt.Println("Before :", p.Name())
// 	p = p.SetName("B")
// 	fmt.Println("After :", p.Name())
// }

func main() {
	p := Person{
		name: "A",
		friends: map[string]int{
			"B": 1,
		},
		test: []int{1, 2, 3},
	}
	fmt.Printf("Before : %#v\n", p)
	p.SetName("B")
	fmt.Printf("After : %#v\n", p)

}
