package candidate

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name        string `json:"name"`
	Document    string `json:"document"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
	Technology  string `json:"technology"`
}
