package models

import "time"

//Product struct for Product
type Product struct {
	ID             string `gorm:"column:id;primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	BaseCurrency   string
	QuoteCurrency  string
	BaseAssetData  string
	QuoteAssetData string
}
