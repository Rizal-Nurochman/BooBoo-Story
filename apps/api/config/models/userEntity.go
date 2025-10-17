package models

import (
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string
	Email     string          `gorm:"unique"`
	Password  string
	Role      string          `gorm:"default:Reader"`
	Avatar    string    `gorm:"default:'https://images.unsplash.com/photo-1628260412297-a3377e45006f?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&q=80&w=1074'"`
	Points    int       `gorm:"default:0"`
	Streak    int       `gorm:"default:0"`
	Level     int       `gorm:"default:1"`
	

	ResetPasswordToken          string    `gorm:"index"`
	ResetPasswordTokenExpiresAt time.Time
	CreatedAt                   time.Time

	Creator *Creator `gorm:"foreignKey:UserID"`

	ProgressReaders []ProgressRead `gorm:"foreignKey:UserID"`
	UserAchievements []UserAchievement `gorm:"foreignKey:UserID"`
	UserInventories  []UserInventory   `gorm:"foreignKey:UserID"`
	UserQuizResults  []UserQuizResult  `gorm:"foreignKey:UserID"`

	// 🔖 Relasi ke bookmark
	BookmarkedStories []*Story `gorm:"many2many:story_bookmarks;"`

	// 👥 Follow system
	Following []*User `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowingID"`
	Followers []*User `gorm:"many2many:user_follows;joinForeignKey:FollowingID;joinReferences:FollowerID"`
}
