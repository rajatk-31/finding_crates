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

//UniqueSpaces struct to store the unique spaces available.
type UniqueSpaces struct {
	CType      string
	TotalCount int32
	FreeCount  int32
	Dimensions Dimensions
}

//SKU defines the sku details
type SKU struct {
	Sku        string
	Quantity   int64
	Dimensions Dimensions
}

//FinalAnswer will tell the final output of the assets assigned for the given inbound order
type FinalAnswer struct {
	Sku    string
	Assets []struct {
		AssetName     string
		TotalQuantity int32
	}
}

//AssignedSpaces gives the final result for the assigned spaces.
type AssignedSpaces struct {
	Success           bool
	SpillPercent      float64
	Qty               int32
	ItemInSingleAsset int32
	FreePerc          float64
	Vol               int64
}
