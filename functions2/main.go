package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func() {
		fmt.Println("Hello Go!")
	}
	f()

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
