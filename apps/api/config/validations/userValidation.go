package validations

import (
	"gorm.io/datatypes"
)

type UserCreate struct {
	Name     string         `json:"name" binding:"required"`
	Email    string         `json:"email" binding:"required,email"`
	Password string         `json:"password" binding:"required,min=8"`
	Role     string         `json:"role" binding:"omitempty,oneof=admin user"`
	Avatar   datatypes.JSON `json:"avatar" binding:"omitempty"`
	Points   int            `json:"points" binding:"omitempty,gte=0"`
	Streak   int            `json:"streak" binding:"omitempty,gte=0"`
	Level    int            `json:"level" binding:"omitempty,gte=1"`
}

type UserRead struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	Avatar    datatypes.JSON `json:"avatar"`
	Points    int            `json:"points"`
	Streak    int            `json:"streak"`
	Level     int            `json:"level"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

type UserUpdate struct {
	Name            string         `json:"name" binding:"omitempty"`
	Email           string         `json:"email" binding:"omitempty,email"`
	Password        string         `json:"password" binding:"omitempty,min=8"`
	ConfirmPassword string         `json:"confirm_password" binding:"omitempty,eqfield=Password"`
	Role            string         `json:"role" binding:"omitempty,oneof=admin user"`
	Avatar          datatypes.JSON `json:"avatar" binding:"omitempty"`
	Points          int            `json:"points" binding:"omitempty,gte=0"`
	Streak          int            `json:"streak" binding:"omitempty,gte=0"`
	Level           int            `json:"level" binding:"omitempty,gte=1"`
}


type UserDelete struct {
	ID uint `json:"id" binding:"required"`
}

