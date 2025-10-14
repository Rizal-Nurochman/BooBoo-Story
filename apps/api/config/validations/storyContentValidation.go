package validations

type StoryContentCreate struct {
	StoryID   uint   `json:"story_id" binding:"required,gt=0"`
	PageOrder int    `json:"page_order" binding:"required,gt=0"`
	Text      string `json:"text" binding:"required,min=1,max=2000"`
	ImageUrl  string `json:"image_url" binding:"omitempty,url"`
	AudioUrl  string `json:"audio_url" binding:"omitempty,url"`
}

type StoryContentUpdate struct {
	PageOrder *int    `json:"page_order" binding:"omitempty,gt=0"`
	Text      *string `json:"text" binding:"omitempty,min=1,max=2000"`
	ImageUrl  *string `json:"image_url" binding:"omitempty,url"`
	AudioUrl  *string `json:"audio_url" binding:"omitempty,url"`
}
