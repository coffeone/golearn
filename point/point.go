package main

import "fmt"

// func inc(p *int) {
// 	*p++
// }
// func main() {
// 	i := 1
// 	inc(&i)
// 	inc(&i)
// 	inc(&i)
// 	fmt.Println(i)
// }

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func main() {
	p := &[3]int{1, 2, 3}
	pow(p)
	fmt.Println(p)
}
