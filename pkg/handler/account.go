package handler

import (
	"net/http"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

func (h *Handler) extend(c *gin.Context) {}

func (h *Handler) details(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	user, err := h.client.GetDetails(c, &pb.GetDetailsRequest{ID: id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
