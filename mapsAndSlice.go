package main

import "log"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	log.Println(numbers)

	log.Println(numbers[6:9])

	names := []string{"John", "Jane", "Mary", "Bob"}

	log.Println(names)
}