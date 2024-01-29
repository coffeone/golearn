package main

import "fmt"

type Base struct {
	Id    int
	Owner string
}

type A struct {
	Base
	Name string
	Area string
}

type B struct {
	Base
	Title  string
	Bodies []string
}

func main() {
	a := A{
		Base: Base{18, "yamada"},
		Name: "yamada",
		Area: "ikebukoro",
	}

	b := B{
		Base:   Base{22, "saito"},
		Title:  "animals",
		Bodies: []string{"A", "B"},
	}
	fmt.Println(a)
	fmt.Println(b)
}
