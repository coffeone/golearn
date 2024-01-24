package main

import "fmt"

// func main() {
// 	s := []int{1, 2, 3}
// 	s = append(s, 4)
// 	fmt.Println(s)
// 	s = append(s, 5, 6, 7)
// 	fmt.Println(s)
// }

//	func main() {
//		s0 := []int{1, 2, 3}
//		s1 := []int{4, 5, 6}
//		s2 := append(s0, s1...)
//		fmt.Println(s2)
//	}
// func main() {
// 	s0 := []int{1, 2, 3}
// 	s1 := []int{4, 5}
// 	s2 := copy(s0, s1)
// 	fmt.Println(s2)
// 	fmt.Println(s0)
// 	fmt.Println(s1)
// }

func main() {
	s := []string{"apple", "banana", "cherry"}

	for i, v := range s {
		fmt.Printf("[%d]=> %s\n", i, v)
	}

}
