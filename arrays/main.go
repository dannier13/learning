package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	var students [3]string
	students[0] = "Lisa"
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student 1: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))

	slices := []int{1, 2, 3}
	fmt.Printf("%v\n", len(slices))
	fmt.Printf("%v\n", cap(slices))

	b := slices[:2]
	fmt.Printf("%v\n", b)

	c := make([]int, 3, 100)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", len(c))
	fmt.Printf("%v\n", cap(c))

	c = append(c, 1, 2)
	c = append(c, []int{3, 4, 5, 6}...)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", len(c))
	fmt.Printf("%v\n", cap(c))

	d := append(c[:3], c[4:]...)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", c)
}
