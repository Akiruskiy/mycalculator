package main

import "fmt"

func main() {
	var mensaje string = "Hola Mundo"
	mensajeFacil := "Hola Mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)
	a := 10.
	var b float64 = 3
	var c int = 10
	d := 3
	fmt.Println(c / d)
	fmt.Println(a)
	fmt.Println(b)
	x := true
	y := false
	fmt.Println(x || y)
	privada()
	Publica()
	imprimirHola()
}

func privada() {
	fmt.Println("esto es una funcion privada")
}

func Publica() {
	fmt.Println("Esto es una funciona publica")
}

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
