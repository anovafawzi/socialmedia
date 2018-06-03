package models

// Relations show the relations object, represent with two emails
type Relations struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"ID" json:"ID"`
	Email1    string `gorm:"not null" form:"email1" json:"email1"`
	Email2    string `gorm:"not null" form:"email2" json:"email2"`
	Friend    bool   `gorm:"not null" form:"friend" json:"friend"`
	Subscribe bool   `gorm:"not null" form:"subscribe" json:"subscribe"`
	Block     bool   `gorm:"not null" form:"block" json:"block"`
}
