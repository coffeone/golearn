package main

import (
	"fmt"
	"zoo/animals"
)

func main() {
	say(animals.CatFood("fish"))
	say(animals.DogFood("cat"))
}

func say(v string) {
	fmt.Println(v)
}
