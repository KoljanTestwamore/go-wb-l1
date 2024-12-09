package main

import "fmt"

type Human struct {
	age int
}

func (h Human) sayAge() {
	fmt.Printf("My age is %v", h.age)
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{age: 27},
	}
	
	action.sayAge()
}