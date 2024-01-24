package main

import "fmt"

func main() {
	var x interface{} = 3
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer: %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type")
	}

	var y interface{} = 7
	switch v := y.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("%#v\n", v)
	}
}
