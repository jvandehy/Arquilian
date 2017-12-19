package main

import (
	"fmt"
)

type spin int
var x spin
var y int
func main() {
	x=42
	y=int(x)
	fmt.Printf("%v%v\n%v%T\n","value of y is ",y,"type of y is ",y)
}
