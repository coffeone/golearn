package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	n := 1
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println(("two"))
	}
}
