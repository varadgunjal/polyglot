package main

import (
	"fmt"
)

func main() {
	var z complex128 = 5 + 12i
	x := 10 == 10
	fmt.Printf("Type : %T, Value : %v\n", z, z)
	fmt.Printf("Type : %T, Value : %v\n", x, x)
}
