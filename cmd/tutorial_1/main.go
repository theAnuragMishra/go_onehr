package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     int
	gallons int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

func (e gasEngine) milesLeft() int {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() int {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() int
}

func canMakeIt(e engine, miles int) {
	if miles <= e.milesLeft() {
		fmt.Println("you can make it there")
	} else {
		fmt.Println("you cant make it there")
	}
}

func main() {
	miles := 20
	gas1 := gasEngine{10, 0}
	electric1 := electricEngine{10, 1}
	canMakeIt(electric1, miles)
	canMakeIt(gas1, miles)
}
