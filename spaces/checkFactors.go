package spacing

import (
	"fmt"
	structs "sample/structs"
)

func unique(strSlice []structs.Product) []structs.UniqueSpaces {
	keys := make(map[string]bool)
	unique := []structs.UniqueSpaces{}
	list := []string{}
	for _, entry := range strSlice {
		// fmt.Println("YHA2", list)
		value := keys[entry.PName]
		// fmt.Println("vl", value)
		if !value {
			keys[entry.PName] = true
			list = append(list, entry.PName)
			x := structs.UniqueSpaces{CType: entry.PName, TotalCount: 1, FreeCount: 0, Dimensions: entry.Dimensions}
			unique = append(unique, x)
			uniqueLen := len(unique)
			if entry.Status == 1 {
				unique[uniqueLen-1].FreeCount += 1
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
			unique[indexofele].TotalCount += 1
			if entry.Status == 1 {
				unique[indexofele].FreeCount += 1
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
	sku := structs.SKU{Sku: "ABC", Quantity: 8, Dimensions: structs.Dimensions{Length: 11, Breadth: 5, Height: 5}}
	fmt.Println("Here is the list", list)
	AllskuObO(list, sku)
	// for _, prod := range p {
	// 	fmt.Println(prod)
	// // }
	// fmt.Println("Hi")
}
