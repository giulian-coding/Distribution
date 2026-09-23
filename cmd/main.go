package main

import (
	"fmt"
	"log"
	"os"

	"github.com/giulian-coding/Distribution/internal/auth"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenManager, err := auth.NewTokenManager(jwtSecret)
	if err != nil {
		log.Fatalf("error while creating the Token Manager: %v", err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("error retrieving hostname: %v", err)
	}

	port := 8080

	token, err := tokenManager.CreateToken(host, port)
	if err != nil {
		log.Fatalf("token coulnd be created: %v", err)
	}
	fmt.Println("created token:", token)

	if err := tokenManager.VerifyToken(token); err != nil {
		log.Fatalf("token verification failed: %v", err)
	}
	fmt.Println("token verification successful")

}
