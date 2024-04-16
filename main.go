package main

import (
	"fmt"

	"github.com/abdulhodiy1020/test_repo/user"
)


func main() {
	var name string
	var age int
	var email string

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Print("Age: ")
	fmt.Scanln(&age)

	fmt.Print("Email: ")
	fmt.Scanln(&email)

	user := User{
		Name:  name,
		Age:   age,
		Email: email,
	}

	errors := user.ValidateUser()

	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("\nUser information is valid.")
	}
}