package spacing

import (
	"fmt"
	"math"
	structs "sample/structs"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//to round ro specified decimal digits
func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*float64(decimals)) / float64(decimals)
}

//Check factors for assigning to single sku
func checkFact(asset structs.UniqueSpaces, sku structs.SKU, lastVol int, lastFreePercent float64, lastSpillPercent float64) structs.AssignedSpaces {
	assetsVol := asset.Dimensions.Length * asset.Dimensions.Breadth * asset.Dimensions.Height
	// totalAvailAssetVol := assetsVol * asset.FreeCount
	assetFreePercent := roundTo((float64(asset.FreeCount) * 100 / float64(asset.TotalCount)), 2)
	skuVol := sku.Dimensions.Length * sku.Dimensions.Breadth * sku.Dimensions.Height
	totalVol := float64(sku.Quantity) * float64(skuVol)

	/**
	 * Checking
	 * -- Total asset Required of one particular type to contain all items
	 * -- Items that can fit in single asset of that particular type
	 * -- Total Spillage in the assets after filling them
	 */

	assetReq := math.Ceil(float64(totalVol) * 1.1 / float64(assetsVol))
	if sku.Quantity == 1 {
		assetReq = math.Ceil(float64(totalVol) / float64(assetsVol))
	}
	item_fit_single_asset := math.Floor(float64(assetsVol) / float64(skuVol))
	if sku.Quantity > 2 {
		item_fit_single_asset = math.Floor(float64(assetsVol) / (float64(skuVol) * 1.1))
	}
	item_in_single_asset := int64(item_fit_single_asset)
	if sku.Quantity <= int64(item_fit_single_asset) {
		item_in_single_asset = sku.Quantity
	}
	spill_in_one := roundTo(((1 - ((float64(item_in_single_asset) * float64(skuVol)) / float64(assetsVol))) * 100), 2)
	if item_in_single_asset > 2 {
		spill_in_one = roundTo(((1 - ((float64(item_in_single_asset) * float64(skuVol) * 1.10) / float64(assetsVol))) * 100), 2)

		// fmt.Println("Je lo asset Req", item_in_single_asset, skuVol, assetsVol, spill_in_one, (1-((float64(item_in_single_asset)*float64(skuVol)*1.10)/float64(assetsVol)))*100)
	}

	reqdQty := math.Ceil(float64(sku.Quantity) / float64(item_in_single_asset))
	total_spill := spill_in_one * reqdQty

	if (int32(assetReq) <= asset.FreeCount) && ((lastSpillPercent > total_spill || lastSpillPercent == -1) || (lastSpillPercent == total_spill && (lastFreePercent > assetFreePercent || (lastFreePercent == assetFreePercent && int32(lastVol) < assetsVol)))) {
		return structs.AssignedSpaces{
			Success:           true,
			SpillPercent:      total_spill,
			Qty:               int32(reqdQty),
			ItemInSingleAsset: int32(item_in_single_asset),
			FreePerc:          assetFreePercent,
			Vol:               int64(assetsVol),
		}
	}

	// if int32(assetReq) <= asset.FreeCount {
	// 	fmt.Println(lastSpillPercent, total_spill)
	// 	if lastSpillPercent > total_spill || lastSpillPercent == -1 {
	// 		fmt.Println("yaha-------1")
	// 		return structs.AssignedSpaces{
	// 			Success:           true,
	// 			SpillPercent:      total_spill,
	// 			Qty:               int32(reqdQty),
	// 			ItemInSingleAsset: int32(item_in_single_asset),
	// 			FreePerc:          assetFreePercent,
	// 			Vol:               int64(assetsVol),
	// 		}
	// 	} else if lastSpillPercent == total_spill && lastFreePercent < assetFreePercent {
	// 		fmt.Println("yaha-------2")
	// 		return structs.AssignedSpaces{
	// 			Success:           true,
	// 			SpillPercent:      total_spill,
	// 			Qty:               int32(reqdQty),
	// 			ItemInSingleAsset: int32(item_in_single_asset),
	// 			FreePerc:          assetFreePercent,
	// 			Vol:               int64(assetsVol),
	// 		}
	// 	} else if lastSpillPercent == total_spill && lastFreePercent == assetFreePercent && int32(lastVol) < assetsVol {
	// 		fmt.Println("yaha-------3")
	// 		return structs.AssignedSpaces{
	// 			Success:           true,
	// 			SpillPercent:      total_spill,
	// 			Qty:               int32(reqdQty),
	// 			ItemInSingleAsset: int32(item_in_single_asset),
	// 			FreePerc:          assetFreePercent,
	// 			Vol:               int64(assetsVol),
	// 		}
	// 	}
	// }
	return structs.AssignedSpaces{Success: false,
		SpillPercent:      0.00,
		Qty:               0,
		ItemInSingleAsset: 0,
		FreePerc:          0.00,
		Vol:               0}

}

//AllskuObo function takes each sku from an array and then assign spaces to each one of them.
func AllskuObO(assets []structs.UniqueSpaces, skus structs.SKU) {
	fmt.Println("Here i Am")
	leastSpillage := "false"
	qty := -1
	lastVol := -1
	lastSpillPercent := -1.00
	lastFreePercent := 100.00
	item_in_single_asset := 0

	for i := 0; i <= len(assets); i++ {

		if i >= len(assets) {
			fmt.Println("Here will be the final answer")
			fmt.Printf(`
				*****************************
				leastSpillage = %s
				lastVol = %d
				qty =  %d
				item_in_single_asset =  %d
				lastSpillPercent = %.2f
				lastFreePercent = %.2f
				*****************************
			`, leastSpillage, lastVol, qty, item_in_single_asset, lastSpillPercent, lastFreePercent)
			// break
		} else {
			checkedFactors := checkFact(assets[i], skus, lastVol, lastFreePercent, lastSpillPercent)
			fmt.Println(checkedFactors)
			if checkedFactors.Success {
				fmt.Println("Ithhe bhi aaya ", assets[i].CType)
				leastSpillage = assets[i].CType
				lastVol = int(checkedFactors.Vol)
				qty = int(checkedFactors.Qty)
				item_in_single_asset = int(checkedFactors.ItemInSingleAsset)
				lastSpillPercent = checkedFactors.SpillPercent
				lastFreePercent = checkedFactors.FreePerc
			}

		}

	}

}
