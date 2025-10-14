package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("mysecret")

	// สร้าง token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
		"name":    "John Doe",
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("JWT Token:", tokenString)

	// สร้างโฟลเดอร์ tokens ถ้ายังไม่มี
	folder := "tokens"
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// เขียน token ลงไฟล์
	filePath := filepath.Join(folder, "jwt.txt")
	err = os.WriteFile(filePath, []byte(tokenString), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JWT token saved to %s\n", filePath)
}
