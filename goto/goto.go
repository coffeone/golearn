package main

import "fmt"

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L:
	fmt.Println("C")
}
