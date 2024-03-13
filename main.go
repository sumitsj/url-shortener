package main

import (
	"fmt"
	"github.com/sumitsj/url-shortener/server"
)

func main() {
	fmt.Println("Application started.")
	err := server.Start()
	if err != nil {
		fmt.Println("Failed to start server.", err)
	}
	defer fmt.Println("Application ended.")
}
