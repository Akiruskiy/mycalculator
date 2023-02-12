package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 8
	fmt.Println(m1)
	fmt.Println("Mapa num 1 es:", m1["a"])
}
