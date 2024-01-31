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
	f.Seek(0, os.SEEK_END)
	f.WriteString("yeah!")
	os.Rename("foo.txt", "bar.txt")

	os.Getwd()
	os.Mkdir("foo", 0775)
	os.MkdirAll("foo/bar/baz", 0775)
	os.Rename("bar.txt", "foo/bar/baz/foo.txt")

	for _, v := range os.Environ() {
		fmt.Println(v)
	}
	os.Clearenv()
	os.Getenv("HOME")
}
