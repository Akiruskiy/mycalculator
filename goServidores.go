package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	chanel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go checkServer(server, chanel)
		}
		time.Sleep(1 * time.Second) // Duerme el programa por un segundo
		fmt.Println(<-chanel)
		i++
	}

	timePassed := time.Since(start)
	fmt.Printf("Execution time %s\n", timePassed)
}

func checkServer(server string, chanel chan string) {
	_, err := http.Get(server)
	if err != nil {
		chanel <- server + "no esta disponible =("
	} else {
		chanel <- server + "Is working fantastic =P"
	}
}
