package validations

type ResetPasswordInput struct {
	Token               	string `json:"token" binding:"required"`
	Password            	string `json:"password" binding:"required,min=8"`
	PasswordConfirmation 	string `json:"password_confirmation" binding:"required"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}