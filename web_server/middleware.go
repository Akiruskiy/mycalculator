package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authetication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				fmt.Fprintln(w, "\nThis is the time now", start)
				fmt.Fprintf(w, "\njajaja es la hora de ahora creo")

				// Aca habia log.Println() para imprimirlo en la consola
			}()
			f(w, r)
		}
	}
}
