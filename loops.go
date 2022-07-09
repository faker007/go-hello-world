package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"cat", "dog", "fish", "horse", "cow"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	animals2 := make(map[string]string)
	animals2["cat"] = "Nya"
	animals2["dog"] = "Fido"

	for _, animals2 := range animals2 {
		log.Println(animals2)
	}

	for animalType, animal2 := range animals2 {
		log.Println(animalType, animal2)
	}

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"John", "Doe", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 23})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 28})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}