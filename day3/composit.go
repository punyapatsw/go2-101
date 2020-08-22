package main

import "fmt"

type OnePlus struct {
	Version string
}

func (o OnePlus) Info() string {
	return fmt.Sprintf("Info : %s", o.Version)
	// return fmt.Sprintf("Info : %s", "test")
}

type Samsung struct {
	OnePlus
	Version string
}

func main() {
	o := OnePlus{
		Version: "7",
	}
	fmt.Println(o.Info())

	s := Samsung{}
	s.Version = "s10"
	fmt.Println(s.Info())

}
