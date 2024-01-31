package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
	t1 := time.Now()

	d := t0.Sub(t1)
	fmt.Println(d)

	t2 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)
	t3 := time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)

	fmt.Println(t2.Before(t3))

}
