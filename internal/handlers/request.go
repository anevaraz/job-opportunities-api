package handlers

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCandidate struct {
	Name        string `json:"name"`
	Document    string `json:"document"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
	Technology  string `json:"technology"`
}

func (r *CreateCandidate) Validate() error {
	if r.Name == "" && r.Document == "" && r.Email == "" && r.Description == "" && r.Education == "" && r.Experience == "" && r.Technology == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Document == "" {
		return errParamIsRequired("document", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Education == "" {
		return errParamIsRequired("education", "string")
	}
	if r.Experience == "" {
		return errParamIsRequired("experience", "string")
	}
	if r.Technology == "" {
		return errParamIsRequired("technology", "string")
	}
	return nil
}
