package main

import "fmt"

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
}
