package main

import "fmt"

type animal interface {
	move() string
}
type dog struct{}
type fish struct{}
type bird struct{}

func (dog) move() string {
	return "I'm a dog and I walk"
}

func (fish) move() string {
	return "I'm a fish and I swim"
}

func (bird) move() string {
	return "I'm a bird and I fly"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	d := dog{}
	f := fish{}
	b := bird{}
	moveAnimal(d)
	moveAnimal(f)
	moveAnimal(b)
}
