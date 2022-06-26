package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string = "Hello, world!"
	var i int = 1

	whatToSay = "Goodbye, cruel world!"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to ", i)
	
	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("The function return", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Soemthing", "else"
}