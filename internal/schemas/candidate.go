package schemas

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Education   string         `json:"education"`
	Expertise   string         `json:"expertise"`
	Technology  string         `json:"technology"`
	Opportunity []*Opportunity `gorm:"many2many:candidates_opportunities;"`
}
