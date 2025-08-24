package models

import "gorm.io/gorm"

// Kreator merepresentasikan profil tambahan untuk User
type Kreator struct {
	gorm.Model
	IgUrl       string
	TiktokUrl   string
	Description string
	Name        string
	UserID      string `gorm:"unique"`

	// Relasi Belongs To (Opsional, tapi baik untuk dimiliki)
	User *User `gorm:"foreignKey:UserID"`
}