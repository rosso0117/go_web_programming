package main

import (
	. "fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// メソッド定義 func (引数名 型(クラス)名), メソッド名、 返り値の型
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	Println("Area of r1: ", r1.area())
	Println("Area of r2: ", r2.area())
	Println("Area of c1: ", c1.area())
	Println("Area of c2: ", c2.area())
}
