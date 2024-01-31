package main

import (
	"fmt"
	"time"
)

// func main() {
// 	t := time.Now()
// 	fmt.Println((t))

// }

func main() {
	// t := time.Date(1998, 1, 15, 10, 14, 23, 0, time.Local)
	// fmt.Println(t)
	// fmt.Println(t.Year())
	// fmt.Println(t.YearDay())
	// fmt.Println(t.Weekday())
	// fmt.Println(time.Hour)

	t := time.Now()
	t = t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t)

}
