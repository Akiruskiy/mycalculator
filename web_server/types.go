package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// definir cuando 
type User struct {
	Name  string 'json:"name"'
	Email string 'json:"email"'
	Phone string 'json:"phone"'
}

type Metadata interface{}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
