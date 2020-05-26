package db

func (d *TargetDBClient) InsertProduct(product *Product) (int, error) {
	insertedEntry := d.db.Create(product)

	return product.ID, insertedEntry.Error
}

func (d *TargetDBClient) InsertMedia(media *Media) (int, error) {
	insertedEntry := d.db.Create(media)

	return media.ID, insertedEntry.Error
}

func (d *TargetDBClient) GroupInsertProduct(products []*Product) error {
	tx := d.db.Begin()
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
