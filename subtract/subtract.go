package main

import "C"

//export subtract
func subtract(a, b C.int) C.int {
	return a - b
}

func main() {}
