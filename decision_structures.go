package main

import "log"

func main() {
	var isTrue bool

	isTrue = false

	if isTrue == true {
		println("isTrue is true")
	} else {
		println("isTrue is false")
	}

	cat := "cat"

	if cat == "cat" {
		println("cat is a cat")
	} else {
		println("cat is not a cat")
	}

	myNum := 100
	isTrue2 := false

	if myNum > 99 && !isTrue2 {
		println("myNum is greater than 100 and isTrue is true")
	} else if myNum < 100 {
		println("myNum is less than 100 or isTrue is false")
	} else if myNum == 101 || isTrue {
		println("myNum is 101 or isTrue is true")
	} else {
		println("myNum is 100 or isTrue is false2")
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
		
	case "dog":
		log.Println("cat is set to dog")
	
	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is something else") 
	}
}