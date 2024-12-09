package main

import "fmt"

type Animal interface {
	Pet()
}

type Cat struct {}

func (c *Cat) Pet() {
	fmt.Print("Meow!")
}

type Dog struct {}

func (d *Dog) feed() {
	fmt.Print("Woof!")
}

type AnimalAdapter struct {
	dog *Dog
}

func (a *AnimalAdapter) Pet() {
	a.dog.feed()
}

func NewAdapter(d *Dog) Animal {
	return &AnimalAdapter{
		dog: d,
	}
}

func PetAnimal(a Animal) {
	a.Pet()
}

func main() {
	cat := Cat{}
	dog := Dog{}


	PetAnimal(&cat)
	PetAnimal(NewAdapter(&dog))


}