package main

import (
	"fmt"
	spacing "sample/spacing"
)

type dimensions struct {
	length  int32
	breadth int32
	height  int32
}

type product struct {
	pName  string
	status int32
	dimensions
}

func appendCrated(c *[]product) {
	// c := *cd
	// Loops to append crates
	for i := 0; i < 5; i++ {
		n := product{pName: "crate", status: 1, dimensions: dimensions{length: 10, breadth: 10, height: 10}}
		*c = append(*c, n)
	}
	for i := 0; i < 5; i++ {
		n := product{pName: "Big", status: 1, dimensions: dimensions{length: 22, breadth: 22, height: 22}}
		*c = append(*c, n)
	}
	for i := 0; i < 5; i++ {
		n := product{pName: "Bigger", status: 1, dimensions: dimensions{length: 50, breadth: 50, height: 50}}
		*c = append(*c, n)
	}
	// return c
}

func main() {
	fmt.Println("This is test file.")
	allProducts := []product{}
	appendCrated(&allProducts)
	fmt.Println(allProducts)
	spacing.EmptySpaces()
}
