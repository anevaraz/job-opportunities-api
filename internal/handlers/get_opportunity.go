package handlers

import (
	"net/http"

	opportunity "github.com/anevaraz/job-opportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpportunityHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opportunity := opportunity.Opportunity{}
	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opportunity not found")
		return
	}

	sendSuccess(ctx, "get-opportunity", opportunity)
}
