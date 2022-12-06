package main

import (
	"fmt"
)

type personInterface interface {
	runPrintName() string
	runPrintAge() int
}

type person struct {
	name string
	age  int
}

func (txt person) runPrintName() string {
	return txt.name
}
func (txt person) runPrintAge() int {
	return txt.age
}

func runTest(p personInterface) {

	fmt.Println(p)
	fmt.Println(p.runPrintName())
	fmt.Println(p.runPrintAge())

}

func main() {

	p1 := person{"myName", 19}
	runTest(p1)

}
