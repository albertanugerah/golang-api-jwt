package dto

type UserUpdateDTO struct {
	ID       uint64  `json:"id" form:"id"`
	Name 	string 	`json:"name" form:"title" binding:"required"`
	Address string 	`json:"address"`
	City 	string 	`json:"city" form:"city"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty" form:"password,omitempty"`
}