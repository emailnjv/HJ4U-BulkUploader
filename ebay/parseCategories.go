package ebay

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

func (ec *EbayClient) loadStoreCategories() error {
	getStoreResponse, err := ec.GetStore()
	if err != nil {
		return err
	}

	var logCategoryID func(utils.GSRCustomCategory)
	logCategoryID = func(customCategories utils.GSRCustomCategory) {
		for _, customCategoryStruct := range customCategories {
			ec.StoreCategoryMap[customCategoryStruct.CategoryID] = customCategoryStruct.Name
			if len(customCategoryStruct.ChildCategory) > 0 {
				logCategoryID(customCategoryStruct.ChildCategory)
			}
		}
	}

	logCategoryID(getStoreResponse.Store.CustomCategories.CustomCategory)

	return nil
}

func (ec *EbayClient) ReturnCategories(getItemResponse utils.GetItemResponse) (string, string, error) {
	mainCategoryStarterString := html.UnescapeString(getItemResponse.Item.PrimaryCategory.CategoryName)
	ampSplit := strings.Split(mainCategoryStarterString, "&")
	specificCategorySplit := strings.Split(strings.TrimSpace(ampSplit[len(ampSplit)-1]), ":")
	mainCategoryString := specificCategorySplit[len(specificCategorySplit)-1]

	storeCategoryID, err := strconv.Atoi(getItemResponse.Item.Storefront.StoreCategoryID)
	if err != nil {
		return "", "", err
	}

	storeCategory, exists := ec.StoreCategoryMap[storeCategoryID]
	if !exists {
		return "", "", fmt.Errorf("Category ID: %v not found in mapping\n", storeCategoryID)
	}

	return mainCategoryString, storeCategory, nil
}
