package main

import (
	"fmt"
	"reflect"
)

type doctor struct {
	number     int
	actorName  string
	companions []string
}

type animal struct {
	name   string
	origin string
}

type animal2 struct {
	name   string `requiredMax:"100"`
	origin string
}

type bird struct {
	animal
	speedKPH float32
	CanFly   bool
}

func main() {
	statePopulations := map[string]int{
		"California": 111,
		"Texas":      222,
		"Florida":    333,
	}
	statePopulations["Georgia"] = 444
	delete(statePopulations, "California")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Florida"])

	pop, ok := statePopulations["George"]
	fmt.Println(pop, ok)

	pop2, ok2 := statePopulations["Georgia"]
	fmt.Println(pop2, ok2)

	_, ok3 := statePopulations["Georgia"]
	fmt.Println(ok3)

	fmt.Println(len(statePopulations))

	aDoctor := doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	aDoctor2 := doctor{
		3,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor2)

	aDoctor3 := struct{ name string }{name: "John Petwee"}
	fmt.Println(aDoctor3)

	b := bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	b2 := bird{
		animal:   animal{name: "Emu", origin: "Australia"},
		speedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b2)

	t := reflect.TypeOf(animal2{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
