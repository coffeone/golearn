package main

import (
	"fmt"
	"time"
)

// func main() {
// 	t := time.Unix(145414, 0)
// 	fmt.Println(t)
// }

func main() {
	// ch := time.Tick(3 * time.Second)
	// for {
	// 	t := <-ch
	// 	fmt.Println(t)
	// }

	ch1 := time.After(5 * time.Second)
	fmt.Println(<-ch1)
}
