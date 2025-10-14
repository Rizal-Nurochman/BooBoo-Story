package models

import (
	"time"
	"gorm.io/datatypes"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string
	Email     string          `gorm:"unique"`
	Password  string
	Role      string          `gorm:"default:Reader"`
	Avatar    datatypes.JSON
	Points    int
	Streak    int
	Level     int
	

	ResetPasswordToken          string    `gorm:"index"`
	ResetPasswordTokenExpiresAt time.Time
	CreatedAt                   time.Time

	Creator *Creator `gorm:"foreignKey:UserID"`

	Stories          []Story           `gorm:"foreignKey:UserID"`
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
