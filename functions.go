package main

import "fmt"

/*
(x, y int) == (x int, y int)
*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
