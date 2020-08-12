package main

import "fmt"

type StudentInfo struct {
	Name, class string
}

// func (student StudentInfo) Info() {
// 	fmt.Println(student.Name, student.class)
// }

func (student StudentInfo) Info() string {
	return fmt.Sprintf("%s, %s", student.Name, student.class)
}

func main() {
	// var m map[string]string
	// m = make(map[string]string)
	// m["pun"] = "Punyapat"
	// fmt.Println(m["pun"])
	// fmt.Println(m["x"])

	var m = map[string]StudentInfo{
		"A0001": StudentInfo{
			"Punyapat",
			"5/5",
		},
	}

	x := StudentInfo{"A", "6/6"}
	m["A0002"] = x
	m["A0003"] = StudentInfo{"B", "4/4"}

	fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Print(k, ":")
	// 	v.Info()
	// }

	for k, v := range m {
		fmt.Println(k, ":", v.Info())
	}
}
