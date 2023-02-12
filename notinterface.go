package main

import "fmt"

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) walk() string {
	return "I'm a dog and I walk"
}

func (fish) swim() string {
	return "I'm a fish and I swim"
}

func (bird) fly() string {
	return "I'm a bird and I fly"
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	d := dog{}
	f := fish{}
	b := bird{}
	moveDog(d)
	moveFish(f)
	moveBird(b)
}
