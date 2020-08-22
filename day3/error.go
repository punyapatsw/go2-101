package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
}

func printErr(err error) {
	fmt.Println(err)
}

func (berr BusinessError) Error() string {
	return "this my my biz error"
}

func main() {
	var err error = errors.New("my error msg")
	printErr(err)
	berr := BusinessError{}
	printErr(berr)
}
