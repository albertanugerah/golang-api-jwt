package dto

type ProfileUpdateDTO struct {
	ID 		uint 	`json:"id" form:"id" binding:"required"`
	Name 	string 	`json:"name" form:"title" binding:"required"`
	Address string 	`json:"address" form:"address" binding:"required"`
	City 	string 	`json:"city" form:"city" binding:"required"`
	UserID	uint 	`json:"user_id,omitempty" form:"user_id"`
}

type ProfileCreateDTO struct {
	Name 	string `json:"name" form:"title" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	City 	string `json:"city" form:"city" binding:"required"`
	UserID	uint 	`json:"user_id,omitempty" form:"user_id"`

}
