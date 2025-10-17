package validations

type CreateProgressInput struct {
	StoryID     uint `json:"story_id" binding:"required"`
	CurrentPage int  `json:"current_page" binding:"required,min=1"`
}

type UpdateProgressInput struct {
	CurrentPage  int  `json:"current_page" binding:"required,min=1"`
	Completed    bool `json:"completed"` 
	EarnedPoints int  `json:"earned_points"` 
}