package entity

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name 	 string 	 `gorm:"type:varchar(100)" json:"name"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	City 	string `gorm:"type:varchar(50)" json:"city"`
	Email    string  `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`

}
