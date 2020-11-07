package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	j := 0
	for ; j < 5; j++ {
		fmt.Println(j)
	}
	fmt.Println("")

	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
	fmt.Println("")

	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("")

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i)
	}
	fmt.Println("")

	s := []int{8, 9, 10}
	fmt.Println(s)
	fmt.Println(s[1])

	for k, v := range s {
		fmt.Println(k, v)
	}

	r := "hello go!"
	for k, v := range r {
		fmt.Println(k, string(v))
	}

	t := "hello go!"
	for _, v := range t {
		fmt.Println(string(v))
	}

	fmt.Println("-----------------------")

}
