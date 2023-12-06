package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func generateToken(userName, password string) string {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15")
	tokenString := fmt.Sprintf("hhm1997%s%s%s%s", userName, password, currentDate, currentTime)
	hashedToken := sha256.Sum256([]byte(tokenString))
	return fmt.Sprintf("%x", hashedToken)
}

func main() {
	userName := "your_username"
	password := "your_password"
	token := generateToken(userName, password)
	fmt.Println("Generated Token:", token)
}
