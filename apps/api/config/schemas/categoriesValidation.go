package validations

type CategoryCreate struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	ParentID *uint  `json:"parent_id,omitempty"`
}

type CategoryUpdate struct {
	Name string `json:"name" validate:"omitempty,min=3,max=100"`
}

type SubCategoryCreate struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	ParentID *uint   `json:"parent_id" validate:"required"`
}