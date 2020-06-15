package db

type CategoryIDMapping map[string]CategoryIDStruct
type CategoryIDStruct struct {
	Name           string
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
func (dbc *DBClient) LevenshteinMatch(input string, categoryName string, targetScore int) (CategoryIDStruct, error) {
	var result CategoryIDStruct
	highestScore := 0

	rows, err := dbc.db.Raw(
		`SELECT ratio, main_cat_id, id from
			      (
							SELECT levenshtein_ratio(name, ?) as ratio, main_cat_id, id from sub_categories WHERE
								main_cat_id = (SELECT id from main_categories WHERE name = ?)
						) ratioTable
					WHERE ratioTable.ratio > ?`,
		input,
		categoryName,
		targetScore,
	).Rows()
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var score int
		var mainCatID int
		var subCatID int

		err = rows.Scan(&score, &mainCatID, &subCatID)
		if err != nil {
			return result, err
		}

		if score > highestScore {
			highestScore = score
			result.MainCategoryID = mainCatID
			result.SubCategoryID = subCatID
		}
	}

	return result, err
}
func (dbc *DBClient) SQLMatchMatch(input string, categoryName string) ([]CategoryIDStruct, error) {
	var result []CategoryIDStruct

	rows, err := dbc.db.Raw(
		`
			SELECT name, main_cat_id, id
			FROM (SELECT * from sub_categories WHERE main_cat_id = (SELECT id from main_categories WHERE name = ?)) fsc
			WHERE LOWER(name) LIKE CONCAT('%', LOWER(?), '%')
		`,
		categoryName,
		input,
	).Rows()
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

		result = append(result, CategoryIDStruct{
			Name: name,
			MainCategoryID: mainCatID,
			SubCategoryID:  subCatID,
		})
	}

	return result, err
}
func (dbc *DBClient) AttemptCategoryMatch(input string, categoryName string, targetScore int) (CategoryIDStruct, error) {
	categoryIDStructs, err := dbc.SQLMatchMatch(input, categoryName)
	if err != nil {
		return CategoryIDStruct{}, err
	}

	categoryIDStructsLength := len(categoryIDStructs)
	switch true {
	case categoryIDStructsLength == 0:
		levenMatch, err := dbc.LevenshteinMatch(input, categoryName, targetScore)
		return levenMatch, err
	case categoryIDStructsLength == 1:
		return categoryIDStructs[0], err
	case categoryIDStructsLength > 1:
		lavenMatch, err := dbc.LevenshteinMatch(input, categoryName, targetScore)
		if err != nil {
			return CategoryIDStruct{}, nil
		}

		for _, idStruct := range categoryIDStructs {
			if idStruct.SubCategoryID == lavenMatch.SubCategoryID {
				return idStruct, err
			}
		}
		return categoryIDStructs[0], err
	default:
		return CategoryIDStruct{}, err
	}
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
