package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("foo.txt")

	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	fi.Size()
	fi.IsDir()

	f.Write([]byte("Hello,world!\n"))
	f.WriteAt([]byte("Golang"), 7)
}
