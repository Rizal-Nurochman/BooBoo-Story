package validations

// Validator untuk membuat story baru
type StoryCreate struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`          // wajib
	Description string `json:"description" binding:"required,min=10"`           // wajib
	CoverImage  string `json:"cover_image" binding:"omitempty,url"`             // opsional
	AgeLevel    string `json:"age_level" binding:"required,oneof=child teen adult"` // wajib, pilih dari opsi
	Difficulty  int    `json:"difficulty" binding:"required,gte=1,lte=5"`       // wajib, rentang 1–5
	CreatorID   uint   `json:"creator_id" binding:"required,gt=0"`              // wajib, harus > 0
}

type StoryUpdate struct {
	Title       *string `json:"title" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description" binding:"omitempty,min=10"`
	CoverImage  *string `json:"cover_image" binding:"omitempty,url"`
	AgeLevel    *string `json:"age_level" binding:"omitempty,oneof=child teen adult"`
	Difficulty  *int    `json:"difficulty" binding:"omitempty,gte=1,lte=5"`
}
