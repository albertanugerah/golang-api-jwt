package dto

type RegisterDTO struct {
	Name 		string 	`json:"name" form:"title" binding:"required"`
	Address 	string 	`json:"address" form:"address" binding:"required"`
	City 		string 	`json:"city" form:"city" binding:"required"`
	Email 		string 		`json:"email" form:"email" binding:"required,email"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}
