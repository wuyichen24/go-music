package main

import (
	"go-music/gin-app/rest"
	"log"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8080"))
}
