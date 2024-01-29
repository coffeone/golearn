package main

import "fmt"

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}
func main() {

	ip := IntPair{1, 2}
	ip.First()
	ip.Last()

	a := Strings{"A", "B", "C"}.Join(",")
	fmt.Println(ip.First())
	fmt.Println(ip.Last())

	fmt.Println(a)
}
