package main

import "fmt"

const (
	a = iota
	_
	b
	c
)

const (
	a2 = iota
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(a2)
}
