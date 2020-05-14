package db

func (d *TargetDBClient) insertProduct(product *Product) (int, error) {
	insertedEntry := d.db.Create(product)

	return product.ID, insertedEntry.Error
}
