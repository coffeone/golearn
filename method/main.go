package main

import "fmt"

type Point struct{ X, Y int }

func (P *Point) Render() {
	fmt.Printf("%d,%d\n", P.X, P.Y)
}
func main() {
	P := &Point{X: 5, Y: 12}
	P.Render()

}
