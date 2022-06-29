package main

import (
	"log"
	"time"
)

var s = "Seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	birthDate time.Time
}

func main() {
	user := User {
		FirstName: "Mirai",
		LastName: "Kim",
		PhoneNumber: "1 555-555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birth date", user.birthDate)
}