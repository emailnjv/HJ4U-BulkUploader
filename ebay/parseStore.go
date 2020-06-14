package ebay

func (ec *EbayClient) loadStoreCategories() error {
	getStoreResponse, err := ec.GetStore()
	if err != nil {
		return err
	}

	for _, customCategoryStruct := range getStoreResponse.Store.CustomCategories.CustomCategory {
		ec.StoreCategoryMap[customCategoryStruct.CategoryID] = customCategoryStruct.Name
	}

	return nil
}
