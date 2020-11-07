package main

import "fmt"

func main() {
	if false {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California": 111,
		"Texas":      222,
		"Florida":    333,
	}

	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 70

	//	if guess < 1 || returnTrue() || guess > 100 {
	//		fmt.Println("The guess must be between 1 and 100!")
	//	}

	//if guess >= 1 && guess <= 100 {
	//	if guess < number {
	// fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("Too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// }

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)
	fmt.Println("")

	switch 3 {
	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4, 6:
		fmt.Println("even")
	default:
		fmt.Println("too big")
	}
	fmt.Println("")

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
	fmt.Println("")

	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
	fmt.Println("")
	fmt.Println("-----------------------------")
}

func returnTrue() bool {
	fmt.Println("returning true!")
	return true
}
