package models

import (
	"gorm.io/gorm"
)

type StoryContent struct {
	ID        uint   `gorm:"primaryKey"`
	StoryID   uint
	PageOrder int
	Text      string
	ImageUrl  string
	AudioUrl  string

	RareWords []StoryContentRareWord `gorm:"foreignKey:StoryContentID"`
}

// 🔧 BeforeCreate hook: otomatis set PageOrder per StoryID
func (sc *StoryContent) BeforeCreate(tx *gorm.DB) (err error) {
	if sc.PageOrder == 0 {
		var maxOrder int
		tx.Model(&StoryContent{}).
			Where("story_id = ?", sc.StoryID).
			Select("COALESCE(MAX(page_order), 0)").
			Scan(&maxOrder)
		sc.PageOrder = maxOrder + 1
	}
	return
}
