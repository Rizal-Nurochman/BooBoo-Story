package models

import "gorm.io/gorm"

// Kreator merepresentasikan profil tambahan untuk User
type Kreator struct {
	gorm.Model
	IgUrl       string
	TiktokUrl   string
	Bio 		string
	Name        string
	UserID      uint	 `gorm:"unique"`

	User *User `gorm:"foreignKey:UserID"`
}