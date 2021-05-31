package dto

type UserUpdateDTO struct {
	ID       uint64  `json:"id" form:"id" binding:"required"`
	Name 	string 	`json:"name" form:"title" binding:"required"`
	Address string 	`json:"address" form:"address" binding:"required"`
	City 	string 	`json:"city" form:"city" binding:"required"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password string  `json:"password,omitempty" form:"password,omitempty" binding:"required"`
}