package entity

type Profile struct {
		ID 		uint `gorm:"primaryKey:auto_increment" json:"id"`
		Name 	string `gorm:"type:varchar(100)" json:"name"`
		Address string `gorm:"type:varchar(255)" json:"address"`
		City 	string `gorm:"type:varchar(50)" json:"city"`
		UserID	uint 	`gorm:"not null" json:"-"`
		User	User	`gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}