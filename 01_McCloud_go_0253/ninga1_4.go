package main

import (
	"fmt"
)

type spin int
var x spin

func main() {
	fmt.Printf("%v%v\n%v%T\n","value is ",x,"type is ",x)
	x = 42
	fmt.Printf("%v\n", x)
}
