package main

import (
	"fmt"
	"log"

	"github.com/giulian-coding/Distribution/internal/auth"
)

func main() {
	host := "test"
	port := 8080

	token, err := auth.CreateToken(host, port)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(token)
}
