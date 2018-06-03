package models

// Connections : struct to get array of emails
type Connections struct {
	Friends []string `gorm:"not null" form:"friends" json:"friends"`
}
