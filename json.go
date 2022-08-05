package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	HairColor string `json:"hairColor"`
	HasDog bool `json:"hasDog"`
}

func main() {
	myJson:= `
[
	{
		"FirstName": "Clark",
		"LastName": "Kent",
		"HairColor": "black",
		"HasDog": true
	},
	{
		"FirstName": "Bruce",
		"LastName": "Wayne",
		"HairColor": "black",
		"HasDog": false
	}
]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling JSON:", err) 
	}

	log.Println("unmarshalled: %v", unmarshalled)

	// write json from a struct

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "green"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false
	
	mySlice = append(mySlice, m2)

	newJson, err := json.Marshal()
}