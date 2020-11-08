package main

import "fmt"

func main() {
	sayMessage("Hello Go!")

	for i := 0; i < 5; i++ {
		sayMessage2("Hello Go!", i)
	}
	fmt.Println("")

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting, name)
	fmt.Println(name)
	fmt.Println("")

	greeting2 := "Hello"
	name2 := "Stacey"
	sayGreeting2(&greeting2, &name2)
	fmt.Println(name2)
	fmt.Println("")

	s := sum(1, 2, 3, 4)
	fmt.Println("The sum is ", s)
	s2 := sum2(1, 2, 3, 4)
	fmt.Println("The sum is ", *s2)
	fmt.Println("")

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	fmt.Println("------------------")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sayGreeting2(greeting2, name2 *string) {
	fmt.Println(*greeting2, *name2)
	*name2 = "Ted"
	fmt.Println(*name2)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	//result := 0
	for _, v := range values {
		result += v
	}
	//return result
	return
}

func sum2(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
