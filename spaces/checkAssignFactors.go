package spacing

import (
	"fmt"
	structs "sample/structs"
)

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

//Check factors for assigning to single sku
func checkFact(asset structs.UniqueSpaces, sku structs.SKU, lastVol int, lastFreePercent float64, lastSpillPercent float64) int {
	assetsVol := asset.Dimensions.Length * asset.Dimensions.Breadth * asset.Dimensions.Height
	totalAvailAssetVol := assetsVol * asset.FreeCount
	assetFreePercent := (asset.FreeCount * 100 / asset.TotalCount)
	skuVol := sku.Dimensions.Length * sku.Dimensions.Breadth * sku.Dimensions.Height
	totalVol := sku.Quantity * skuVol

 		/**
         * Checking 
         * -- Total asset Required of one particular type to contain all items
         * -- Items that can fit in single asset of that particular type
         * -- Total Spillage in the assets after filling them
         */

	 assetReq := 
	 item_fit_single_asset := Math.floor(sku.qty > 2 ? assetsVol / (skuVol * 1.10) : assetsVol / skuVol);
	 item_in_single_asset := sku.qty <= item_fit_single_asset ? sku.qty : item_fit_single_asset;
	 spill_in_one := (item_in_single_asset > 2 ? (1 - ((item_in_single_asset * skuVol * 1.10) / assetsVol)) * 100 : (1 - ((item_in_single_asset * skuVol) / assetsVol)) * 100).toFixed(2)
}


//AllskuObo function takes each sku from an array and then assign spaces to each one of them.
func AllskuObO(assets []structs.UniqueSpaces, skus structs.SKU) {
	fmt.Println("Here i Am")
	leastSpillage := false
	qty := -1
	lastVol := -1
	lastSpillPercent := -1.00
	lastFreePercent := 100.00
	item_in_single_asset := 0

	for i := 1; i < len(assets); i++ {
		if i >= len(assets) {
			fmt.Println("Here will be the final answer")
			// break
		} else {
			checkedFactors := checkFact(assets[i], skus, lastVol, lastFreePercent, lastSpillPercent)

		}

	}

}
