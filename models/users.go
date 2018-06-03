package models

// Users : the object of a  user
type Users struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"ID" json:"ID"`
	Email     string `gorm:"not null" form:"email" json:"email"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}
