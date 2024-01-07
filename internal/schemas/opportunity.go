package schemas

import "gorm.io/gorm"

type Opportunity struct {
	gorm.Model
	Title        string       `json:"title"`
	Company      string       `json:"company"`
	Description  string       `json:"description"`
	Requirements string       `json:"requirements"`
	Benefits     string       `json:"benefits"`
	Salary       int64        `json:"salary"`
	Remote       bool         `json:"remote"`
	Expertise    string       `json:"expertise"`
	Technology   string       `json:"technology"`
	Candidate    []*Candidate `gorm:"many2many:candidates_opportunities;"`
}
