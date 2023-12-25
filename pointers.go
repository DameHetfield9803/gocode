package main

import "fmt"

func main() {
	var i int = 10
	var p *int = &i
	fmt.Println("i=", i)
	fmt.Println("p=", p)
	fmt.Printf("Value of i is stored in memory address %p.\n", p)
}

// Damien
