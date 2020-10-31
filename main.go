package main

import "fmt"

var i int = 1

var (
	name    = "Dan"
	surname = "Safarik"
	age     = 29
)

func main() {
	var j string
	j = fmt.Sprint(age)
	fmt.Printf("%v, %T", j, j)
}
