package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("lenth=%d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
