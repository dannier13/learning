package main

import "fmt"

var i int = 1

var (
	name    = "Dan"
	surname = "Safarik"
	age     = 29
)

func main() {
	n := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
}
