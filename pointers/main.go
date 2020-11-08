package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Println("%v %v %v\n", x, y, z)

	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println(ms)
	fmt.Println((*ms).foo)

	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
