package handlers

import (
	"net/http"

	candidate "github.com/anevaraz/job-opportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func GetCandidateHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	candidate := candidate.Candidate{}
	if err := db.First(&candidate, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "candidate not found")
		return
	}

	sendSuccess(ctx, "get-candidate", candidate)
}
