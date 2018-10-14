package main

import (
	"fmt"
	"log"
	"os"
	"projects/simple_api/app"
)

func main() {
	fmt.Println("Hello, world!")
	addr := os.Getenv("API_ADDR")
	port := os.Getenv("API_PORT")
	if addr == "" || port == "" {
		log.Fatal("API_PORT or API_ADDR not set!")
		os.Exit(1)
	}

	myapp := app.App{}
	myapp.Run(addr, port)
}
