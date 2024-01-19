package models

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}