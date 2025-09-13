package models

import "time"

// UserInventory adalah join table antara User dan ShopItem
type UserInventory struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	ItemID     uint
	AcquiredAt time.Time

	// Relasi Belongs To
	User User     `gorm:"foreignKey:UserID"`
	Item ShopItem `gorm:"foreignKey:ItemID"`
}