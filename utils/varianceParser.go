package utils

import (
	"fmt"
	"strconv"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
)

type VarianceParser struct{}

type ItemAspects [][]GIGRLocalizedAspect

func (vp *VarianceParser) HandleVariances(productID int, varianceResponse []GIGRItems) ([]*db.ProductAtt, error) {
	var varianceArr ItemAspects
	var result []*db.ProductAtt

	// Create array of aspect arrays
	for _, item := range varianceResponse {
		varianceArr = append(varianceArr, item.LocalizedAspects)
	}

	// Get the unique aspects
	uniqueAspects, err := vp.returnDifferentAspects(varianceArr)
	if err != nil {
		return result, err
	}

	// Iterate over unique aspects generating ready to insert DB product attribute structs
	for itemIndex, uniqueAspect := range uniqueAspects {

		// Get the price
		priceString := varianceResponse[itemIndex].Price.Value
		if priceString == "" {
			return result, fmt.Errorf("Price not found for item:\n%#v", varianceResponse[itemIndex])
		}
		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			return result, err
		}

		// Create and append db.ProductAtt to result
		result = append(result, &db.ProductAtt{
			PID:    productID,
			AKey:   uniqueAspect.Name,
			AValue: uniqueAspect.Value,
			Price: price,
		})
	}
	
	return result, nil
}

func (vp *VarianceParser) returnDifferentAspects(itemAspectGroups ItemAspects) ([]GIGRLocalizedAspect, error) {
	if !vp.equalLengthCheck(itemAspectGroups...) {
		return []GIGRLocalizedAspect{}, fmt.Errorf("inequal length aspects")
	} else {
		return vp.findDifferentLocalizedAspects(itemAspectGroups...), nil
	}
}
func (vp *VarianceParser) equalLengthCheck(localizedArr ...[]GIGRLocalizedAspect) bool {
	if len(localizedArr) > 0 {
		baseLength := len(localizedArr[0])
		for _, aspectCollection := range localizedArr {
			if len(aspectCollection) != baseLength {
				return false
			}
		}
		return true
	}
	return false
}
func (vp *VarianceParser) findDifferentLocalizedAspects(itemAspectArrays ...[]GIGRLocalizedAspect) []GIGRLocalizedAspect {
	itemAspectMap := make(map[string]string)
	var result []GIGRLocalizedAspect

	// range over item arrays
	for itemIndex, itemAspectCollection := range itemAspectArrays {
		// range over an items aspect arrays
		for _, indAspect := range itemAspectCollection {

			// If on the first iteration
			if itemIndex == 0 {

				// Assign the initial mapping
				itemAspectMap[indAspect.Name] = indAspect.Value
			} else {

				// Check if Aspect name previously inserted
				categoryStruct, found := itemAspectMap[indAspect.Name]
				if !found {
					fmt.Printf("Aspect name not found in mapping inside compareEqual\n%#v", categoryStruct)
				}

				// If value is different
				if itemAspectMap[indAspect.Name] != indAspect.Value && indAspect.Name == "Amount" {
					if len(result) == 0 {

						// If first change noticed, add in the initial value
						result = append(result, GIGRLocalizedAspect{
							Name:  indAspect.Name,
							Type:  indAspect.Type,
							Value: itemAspectMap[indAspect.Name],
						})
					}

					// Add the difference
					result = append(result, indAspect)
				}
			}
		}
	}
	return result
}
