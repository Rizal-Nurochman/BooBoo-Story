package models

import "time"

// ShopItem merepresentasikan item di toko
type ShopItem struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Type      string
	Price     int
	AssetUrl  string
	CreatedAt time.Time

	// Relasi Has Many
	UserInventories []UserInventory `gorm:"foreignKey:ItemID"`
}