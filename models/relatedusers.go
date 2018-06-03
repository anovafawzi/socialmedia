package models

// RelatedUsers : struct of an existing related users
type RelatedUsers struct {
	Requestor string `gorm:"not null" form:"requestor" json:"requestor"`
	Target    string `gorm:"not null" form:"target" json:"target"`
}
