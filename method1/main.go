package main

import "fmt"

// import "math"

// type IntPoint struct{ X, Y int }
// type FloatPoint struct{ X, Y float64 }

// func (P *IntPoint) Distance(dp *IntPoint) float64 {
// 	x, y := P.X-dp.X, P.Y-dp.Y
// 	return math.Sqrt(float64(x*x, y*y))
// }

// func (P *FloatPoint) Distance(dp *FloatPoint) float64 {
// 	x, y := P.X-dp.X, P.Y-dp.Y
// 	return math.Sqrt(float64(x*x, y*y))
// }

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}
func main() {
	c := MyInt(4).Plus(2)
	fmt.Println(c)
}
