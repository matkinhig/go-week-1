package main

import ( 
	"fmt"
	"github.com/matkinhig/golang-course/week-1/1/handle"
)

func main() {
	var t handle.Tester()
	t = f
	fmt.Printf("Interface type %T value %v\n", t, t)
	t.Test()
}