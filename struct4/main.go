package main

import (
	"fmt"

	"foo.go/foo"
)

func main() {
	t := &foo.T{}
	t.Method1()
	fmt.Println(t)
}
