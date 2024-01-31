package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 32, "最大値")
	flag.StringVar(&msg, "m", "", "メッセージ")
	flag.BoolVar(&x, "x", false, "拡大オプション")
	flag.Parse()

	fmt.Println("処理数の最大値＝", max)
	fmt.Println("処理メッセージ＝", msg)
	fmt.Println("拡大オプション＝", x)

}
