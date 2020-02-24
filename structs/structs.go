package structs

//Dimensions of the product
type Dimensions struct {
	Length  int32
	Breadth int32
	Height  int32
}

//Product or type of crates
type Product struct {
	PName  string
	Status int32
	Dimensions
}
