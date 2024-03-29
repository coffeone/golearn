package main

import "fmt"

type Point struct {
	X, Y int
}

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nill>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return str
}
func main() {

	ps := Points{}
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	fmt.Println(ps.ToString())
}
