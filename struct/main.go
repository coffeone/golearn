package main

import "fmt"

type Person struct {
	ID   uint
	name string
	部署   string
}

func main() {
	p := Person{ID: 12, name: "山田", 部署: "開発部"}
	fmt.Println(p)
}
