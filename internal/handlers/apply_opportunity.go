package handlers

import (
	"net/http"

	candidate_opportunity "github.com/anevaraz/job-opportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ApplyOpportunityHandler(ctx *gin.Context) {
	request := ApplyOpportunity{}

	ctx.BindJSON(&request)

	if err := request.ValidateApplyOpportunity(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	applyOpportunity := candidate_opportunity.CandidatesOpportunities{
		CandidateId:   request.CandidateId,
		OpportunityId: request.OpportunityId,
	}

	if err := db.Create(&applyOpportunity).Error; err != nil {
		logger.Errorf("error apply opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error apply opportunity on database")
		return
	}

	sendSuccess(ctx, "apply-opportunity", applyOpportunity)
}
