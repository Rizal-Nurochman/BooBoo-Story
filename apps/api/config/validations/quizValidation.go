package validations

type CreateAnswerInput struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"is_correct"`
}

type CreateQuestionInput struct {
	Question string              `json:"question" binding:"required"`
	Answers  []CreateAnswerInput `json:"answers" binding:"required,min=1,dive"`
}

type CreateQuizInput struct {
	StoryID    uint                  `json:"story_id" binding:"required"`
	Difficulty int                   `json:"difficulty" binding:"required,min=1,max=5"`
	Questions  []CreateQuestionInput `json:"questions" binding:"required,min=1,dive"`
}