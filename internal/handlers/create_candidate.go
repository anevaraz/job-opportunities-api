package handlers

import (
	"net/http"

	candidate "github.com/anevaraz/job-opportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCandidateHandler(ctx *gin.Context) {
	request := CreateCandidate{}

	ctx.BindJSON(&request)

	if err := request.ValidateCandidate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newCandidate := candidate.Candidate{
		Name:       request.Name,
		Email:      request.Email,
		Education:  request.Education,
		Expertise:  request.Expertise,
		Technology: request.Technology,
	}

	if err := db.Create(&newCandidate).Error; err != nil {
		logger.Errorf("error creating candidate: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating candidate on database")
		return
	}

	sendSuccess(ctx, "create-candidate", newCandidate)
}
