package entity

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Email    string  `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`

}