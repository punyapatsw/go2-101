package main

import "fmt"

type Student struct {
	Name string
	ID   int
}

func (s Student) String() string {
	return fmt.Sprintf("Hello %s, ID : %d", s.Name, s.ID)
}

func main() {
	s := Student{ID: 10, Name: "Punyapat"}
	fmt.Println(s)
}
