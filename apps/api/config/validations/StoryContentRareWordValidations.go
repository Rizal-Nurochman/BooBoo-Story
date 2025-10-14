package validations

// StoryContentRareWordCreate digunakan untuk validasi pembuatan kata langka
type StoryContentRareWordCreate struct {
	Title           string `json:"title" binding:"required,min=2,max=100"`
	Description     string `json:"description" binding:"omitempty,max=500"`
	StoryContentID  uint   `json:"story_content_id" binding:"required,gt=0"`
}

// StoryContentRareWordUpdate digunakan untuk validasi pembaruan kata langka
type StoryContentRareWordUpdate struct {
	Title       *string `json:"title" binding:"omitempty,min=2,max=100"`
	Description *string `json:"description" binding:"omitempty,max=500"`
}
