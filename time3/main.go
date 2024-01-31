package main

import (
	"fmt"
	"time"
)

// func main() {
// 	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)

// 	t1 := t0.AddDate(1, 0, 0)

// 	fmt.Println(t1)

// 	t3 := t0.Format("1020/01/01")
// 	fmt.Println(t3)
// }

// 時刻をローカルタイムに変換する
func main() {
	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	jst := t.Local()
	fmt.Println(jst)

}
