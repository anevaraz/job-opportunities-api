package handlers

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCandidate struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Education  string `json:"education"`
	Expertise  string `json:"expertise"`
	Technology string `json:"technology"`
}

type CreateOpportunity struct {
	Title        string `json:"title"`
	Company      string `json:"company"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Benefits     string `json:"benefits"`
	Salary       int64  `json:"salary"`
	Remote       *bool  `json:"remote"`
	Expertise    string `json:"expertise"`
	Technology   string `json:"technology"`
}

type ApplyOpportunity struct {
	CandidateId   int `json:"candidateId"`
	OpportunityId int `json:"OpportunityId"`
}

func (r *CreateCandidate) ValidateCandidate() error {
	if r.Name == "" && r.Email == "" && r.Education == "" && r.Expertise == "" && r.Technology == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Education == "" {
		return errParamIsRequired("education", "string")
	}
	if r.Expertise == "" {
		return errParamIsRequired("expertise", "string")
	}
	if r.Technology == "" {
		return errParamIsRequired("technology", "string")
	}
	return nil
}

func (r *CreateOpportunity) ValidateOpportunity() error {
	if r.Title == "" && r.Company == "" && r.Description == "" && r.Requirements == "" && r.Benefits == "" && r.Salary <= 0 && r.Remote == nil && r.Expertise == "" && r.Technology == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Requirements == "" {
		return errParamIsRequired("requirements", "string")
	}
	if r.Benefits == "" {
		return errParamIsRequired("benefits", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "number")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}
	if r.Expertise == "" {
		return errParamIsRequired("expertise", "string")
	}
	if r.Technology == "" {
		return errParamIsRequired("technology", "string")
	}
	return nil
}

func (r *ApplyOpportunity) ValidateApplyOpportunity() error {
	if r.CandidateId <= 0 && r.OpportunityId <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.CandidateId <= 0 {
		return errParamIsRequired("candidateId", "number")
	}
	if r.OpportunityId <= 0 {
		return errParamIsRequired("opportunityId", "number")
	}
	return nil
}
