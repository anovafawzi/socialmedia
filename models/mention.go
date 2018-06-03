package models

// Mention : struct to get users to be mentioned
type Mention struct {
	Sender string `gorm:"not null" form:"sender" json:"sender"`
	Text   string `gorm:"not null" form:"text" json:"text"`
}
