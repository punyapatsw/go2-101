package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	input := strings.ReplaceAll(s, ",", "")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ToLower(input)
	split := strings.Fields(input)
	result := make(map[string]int)
	for _, x := range split {
		result[x] = result[x] + 1
	}
	return result
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}
