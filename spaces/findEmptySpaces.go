package spacing

import (
	"fmt"
	structs "sample/structs"
)

//EmptySpaces Exported function
func EmptySpaces(p []structs.Product) {
	for _, prod := range p {
		fmt.Println(prod)
	}
	fmt.Println("Hi")
}
