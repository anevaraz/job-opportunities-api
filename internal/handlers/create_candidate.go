package handlers

import (
	"net/http"

	candidate "github.com/anevaraz/go-jobs-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	request := CreateCandidate{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newCandidate := candidate.Candidate{
		Name:        request.Name,
		Document:    request.Document,
		Email:       request.Email,
		Description: request.Description,
		Education:   request.Education,
		Experience:  request.Experience,
		Technology:  request.Technology,
	}

	if err := db.Create(&newCandidate).Error; err != nil {
		logger.Errorf("error creating candidate: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating candidate on database")
		return
	}

	sendSuccess(ctx, "create-candidate", newCandidate)
}
