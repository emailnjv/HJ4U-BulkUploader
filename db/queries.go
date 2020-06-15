package db

type CategoryIDMapping map[string]CategoryIDStruct
type CategoryIDStruct struct {
	MainCategoryID int
	SubCategoryID  int
}

func (dbc *DBClient) InsertProduct(product *Products) (int, error) {
	insertedEntry := dbc.db.Create(product)

	return product.ID, insertedEntry.Error
}

func (dbc *DBClient) InsertProductAtt(productAtt *ProductAtt) (int, error) {
	insertedEntry := dbc.db.Create(productAtt)

	return productAtt.ID, insertedEntry.Error
}

func (dbc *DBClient) InsertMedia(media *Media) (int, error) {
	insertedEntry := dbc.db.Create(media)

	return media.ID, insertedEntry.Error
}

func (dbc *DBClient) GroupInsertProduct(products []*Products) error {
	tx := dbc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, product := range products {
		if err := tx.Create(product).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
func (dbc *DBClient) GroupInsertProductAtt(productAttributes []*ProductAtt) error {
	tx := dbc.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, productAttribute := range productAttributes {
		if err := tx.Create(productAttribute).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (dbc *DBClient) LevenshteinRatio(targetWord1 string, targetWord2 string) (int, error) {
	var functionResult int
	row := dbc.db.Raw("SELECT levenshtein_ratio(?,?)", targetWord1, targetWord2).Row()
	err := row.Scan(&functionResult)

	return functionResult, err
}

func (dbc *DBClient) MainCatsSubCats(mainCategoriesName string) (CategoryIDMapping, error) {
	result := make(CategoryIDMapping)

	rows, err := dbc.db.Raw("SELECT name, main_cat_id, id from sub_categories WHERE main_cat_id = (SELECT id from main_categories WHERE name = ?)", mainCategoriesName).Rows()
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var mainCatID int
		var subCatID int

		err = rows.Scan(&name, &mainCatID, &subCatID)
		if err != nil {
			return result, err
		}
		result[name] = CategoryIDStruct{
			MainCategoryID: mainCatID,
			SubCategoryID:  subCatID,
		}
	}

	return result, err
}
