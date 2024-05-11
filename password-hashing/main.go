package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	var err error
	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password)
	fmt.Println("Password :", password)
	fmt.Println("Hash     :", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match    :", match)
}
