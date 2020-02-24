package spacing

import (
	"fmt"
	structs "sample/structs"
)

//UniqueSpaces struct to store the unique spaces available.
type UniqueSpaces struct {
	cType      string
	totalCount int32
	freeCount  int32
	dimensions structs.Dimensions
}

//Unique --> to find unique elements in a slice
func unique(strSlice []structs.Product) []UniqueSpaces {
	keys := make(map[string]bool)
	unique := []UniqueSpaces{}
	list := []string{}
	for _, entry := range strSlice {
		// fmt.Println("YHA2", list)
		value := keys[entry.PName]
		// fmt.Println("vl", value)
		if !value {
			keys[entry.PName] = true
			list = append(list, entry.PName)
			x := UniqueSpaces{cType: entry.PName, totalCount: 1, freeCount: 0, dimensions: entry.Dimensions}
			unique = append(unique, x)
			uniqueLen := len(unique)
			if entry.Status == 1 {
				unique[uniqueLen-1].freeCount += 1
			}
			// fmt.Println("itthe ", unique)

		} else {
			// fmt.Println("itthe2 ", unique)
			indexofele := -1
			// fmt.Println("YHA1")
			for i, v := range list {

				// fmt.Println("YHA13")
				if v == entry.PName {
					indexofele = i
				}
			}

			// fmt.Println("YHA1xx", unique, "-==-=-=-=", indexofele)
			unique[indexofele].totalCount += 1
			if entry.Status == 1 {
				unique[indexofele].freeCount += 1
			}
		}
	}
	// fmt.Println(unique)
	return unique
}

//EmptySpaces Exported function
func EmptySpaces(p []structs.Product) {
	// keys := []struct{
	// 	cType string
	// 	totalCount int32
	// 	freeCount int32
	// 	dimensions structs.Dimensions
	// }
	list := unique(p)
	fmt.Println("Here is the list", list)
	// for _, prod := range p {
	// 	fmt.Println(prod)
	// // }
	// fmt.Println("Hi")
}
