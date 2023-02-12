package main

import (
	"fmt"
	"io"
	"net/http"
)

type WebWriter struct{}

func (WebWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	answer, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	w := WebWriter{}
	io.Copy(w, answer.Body)
}
