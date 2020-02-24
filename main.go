package main

import (
	"fmt"
	spacing "sample/spacing"
	structs "sample/structs"
)

func appendCrated(c *[]structs.Product) {
	// c := *cd
	// Loops to append crates
	for i := 0; i < 5; i++ {
		n := structs.Product{PName: "crate", Status: 1, Dimensions: structs.Dimensions{Length: 10, Breadth: 10, Height: 10}}
		*c = append(*c, n)
	}
	for i := 0; i < 5; i++ {
		n := structs.Product{PName: "Big", Status: 1, Dimensions: structs.Dimensions{Length: 22, Breadth: 22, Height: 22}}
		*c = append(*c, n)
	}
	for i := 0; i < 5; i++ {
		n := structs.Product{PName: "Bigger", Status: 1, Dimensions: structs.Dimensions{Length: 50, Breadth: 50, Height: 50}}
		*c = append(*c, n)
	}
	// return c
}

func main() {
	fmt.Println("This is test file.")
	allProducts := []structs.Product{}
	appendCrated(&allProducts)
	fmt.Println(allProducts)
	spacing.EmptySpaces(allProducts)
}
