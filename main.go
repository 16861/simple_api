package main

import (
	"fmt"
	"projects/simple_api/app"
)

func main() {
	fmt.Println("Hello, world!")

	myapp := app.App{}
	myapp.Run()
}
