package main

import (
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}
func main(){

}

func Validate(req UserCreateRequest) string {

	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return "invalid request"
	}
	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}

	return ""

}
