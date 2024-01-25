package main

import "fmt"

// func main() {
// 	m := make(map[int]string)
// 	m[1] = "China"
// 	m[2] = "us"
// 	m[3] = "Japan"
// 	fmt.Println(m)
// }

// func main() {
// 	m := map[int]string{
// 		1: "Apple",
// 		2: "Banana",
// 		3: "Cherry",
// 	}
// 	for k, v := range m {
// 		fmt.Printf("%d-->> %s\n", k, v)
// 	}
// }

// func main() {
// 	m := map[string]int{
// 		"apple":  88,
// 		"Banana": 99,
// 		"Cherry": 46,
// 	}

// 	m["orange"] = 77
// 	m["lemon"] = 54
// 	m["pear"] = 98
// 	for k, v := range m {
// 		fmt.Printf("%s--> %d\n", k, v)
// 	}
// }

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "c"}
	delete(m, 2)
	fmt.Println(m)
}
