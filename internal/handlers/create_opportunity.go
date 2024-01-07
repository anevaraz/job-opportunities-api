package handlers

import (
	"net/http"

	opportunity "github.com/anevaraz/job-opportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(ctx *gin.Context) {
	request := CreateOpportunity{}

	ctx.BindJSON(&request)

	if err := request.ValidateOpportunity(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newOpportunity := opportunity.Opportunity{
		Title:        request.Title,
		Company:      request.Company,
		Description:  request.Description,
		Requirements: request.Requirements,
		Benefits:     request.Benefits,
		Salary:       request.Salary,
		Remote:       *request.Remote,
		Expertise:    request.Expertise,
		Technology:   request.Technology,
	}

	if err := db.Create(&newOpportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opportunity on database")
		return
	}

	sendSuccess(ctx, "create-opportunity", newOpportunity)
}
