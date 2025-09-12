package main

import "fmt"

func main() {
	var name string
	var classe string
	fmt.Print("Enter your name and classe: ")
	fmt.Scan(&name, &classe) // Reading input separated by space
	fmt.Printf("Hello %s, you are an %s.\n", name, classe)
}
