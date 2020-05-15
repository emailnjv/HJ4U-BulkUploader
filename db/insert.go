package db

func (d *TargetDBClient) InsertProduct(product *Product) (int, error) {
	insertedEntry := d.db.Create(product)

	return product.ID, insertedEntry.Error
}

func (d *TargetDBClient) InsertMedia(media *Media) (int, error) {
	insertedEntry := d.db.Create(media)

	return media.ID, insertedEntry.Error
}