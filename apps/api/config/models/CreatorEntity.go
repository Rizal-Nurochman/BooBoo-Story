package models

import (
	"gorm.io/gorm"
)

type Creator struct {
	gorm.Model
	Name         string
	Bio          string
	ProfileImage string
	BannerImage  string
	Occupation   string
	Location     string

	IgUrl        string
	TiktokUrl    string
	WebsiteURL   string
	YoutubeURL   string
	LinkedInURL  string
	Portfolio    string
	IsVerified   bool			`gorm:"default:false"`
	Status       string 	`gorm:"default:'active'"`

	UserID       uint     `gorm:"unique"`
	User         *User    `gorm:"foreignKey:UserID"`

	Stories      []Story  `gorm:"foreignKey:CreatorID"`
}
