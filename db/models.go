package db

import (
	"time"
)

type Products struct {
	ID             int        `gorm:"column:id;type:int(11);AUTO_INCREMENT;PRIMARY_KEY`
	Name           string     `gorm:"column:name;type:text"`
	Description    string     `gorm:"column:descr;type:text"`
	Price          float64    `gorm:"type:float"`
	Featured       int        `gorm:"type:tinyint(1)"`
	Discount_price int        `gorm:"type:int(11)"`
	Date           *time.Time `gorm:"type:timestamp;NOT NULL"`
	Main_cat       int        `gorm:"type:int(11)"`
	Sub_cat        int        `gorm:"type:int(11)"`
	Qty            int        `gorm:"type:int(11)"`
	Sku            string     `gorm:"type:varchar(200);NOT NULL"`
	Upc            string     `gorm:"type:varchar(200);NOT NULL"`
	Product_type   string     `gorm:"type:varchar(200);NOT NULL"`
}

type ProductAtt struct {
	ID     int     `gorm:"column:id;type:int(11);AUTO_INCREMENT;PRIMARY_KEY`
	PID    int     `gorm:"column:p_id;type:int(11)`
	AKey   string  `gorm:"column:a_key;type:text"`
	AValue string  `gorm:"column:a_value;type:text"`
	Price  float64 `gorm:"type:int(11)`
}

type Media struct {
	ID            int        `gorm:"column:id;type:int(11);AUTO_INCREMENT;PRIMARY_KEY`
	PageID        int        `gorm:"column:page_id;type:int(11)`
	MediaType     string     `gorm:"column:media_type;type:varchar(11)`
	PageType      string     `gorm:"column:page_type;type:varchar(200)`
	MediaFileName string     `gorm:"column:media_file_name;type:text"`
	MediaThumb    string     `gorm:"column:media_thumb;type:text;NOT_NULL"`
	Date          *time.Time `gorm:"type:DEFAULT;CURRENT_TIMESTAMP;NOT NULL"`
	Author        int        `gorm:"type:int(11)`
	OrderID       int        `gorm:"column:order_id;type:int(11)`
}

type SubCategories struct {
	ID            int        `gorm:"column:id;type:int(11);AUTO_INCREMENT;PRIMARY_KEY`
	MainCatID     int        `gorm:"column:main_cat_id;type:varchar(255)`
	Name          string     `gorm:"column:name;type:varchar(255)`
}

type levenshteinRatioResult struct {
	levenshteinRatio  int
}
