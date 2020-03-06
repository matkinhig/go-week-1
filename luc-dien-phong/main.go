package main

import "fmt"

// Point cmt
type Point struct {
	X int
	y int
}

// FuncPointer cmt
func FuncPointer() {
	p1 := Point{9, 10}
	fmt.Println(p1)
	p2 := &p1
	p2.X = -9
	fmt.Println(p1)
}

func main() {
	fmt.Println("Pls run file: homeworkday1_test.go")
	FuncPointer()
}
