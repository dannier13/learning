package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("middle") // executes last
	fmt.Println("end")
	fmt.Println("")

	fmt.Println("----------------------")
}
