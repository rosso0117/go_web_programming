package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名フィールド
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Println("Name: ", mark.name)
	fmt.Println("Age: ", mark.age)
	fmt.Println("Weight: ", mark.weight)
	fmt.Println("Speciality: ", mark.speciality)
}
