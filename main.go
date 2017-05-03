package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	for {
		fmt.Print("Enter string to encrypt: ")
		var input string
		fmt.Scanln(&input)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
		if err != nil {
			fmt.Errorf("Error while hashing: %s", err)
			os.Exit(1)
		}

		fmt.Printf("Hashed: %q", hashedPassword)
		fmt.Println("")
	}
}
