package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)  // te duvuelve 25
	fmt.Println(&x) // te devuelve el puntero del valor, la direccion de memoria de la variable
	changeValue(x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}

func changeValue(a int) {
	fmt.Println(&a)
	a = 36
}
