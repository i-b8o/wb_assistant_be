package handler

import (
	"net/http"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

// @Summary update account
// @Security ApiKeyAuth
// @Tags account
// @Description update account
// @ID update-account
// @Accept json
// @Produce json
// @Param input body pb.UpdateRequest true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /account/update [post]
func (h *Handler) update(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}

	input := &pb.UpdateRequest{}
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" && input.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.ID = id
	_, err = h.authClient.Update(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary account details
// @Security ApiKeyAuth
// @Tags account
// @Description account details
// @ID details-account
// @Accept json
// @Produce json
// @Success 200 {object} pb.User 1
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /account/details [get]
func (h *Handler) details(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}
	user, err := h.authClient.GetDetails(c, &pb.GetDetailsRequest{ID: id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary update actions
// @Security ApiKeyAuth
// @Tags account
// @Description update actions
// @ID update-actions
// @Accept json
// @Produce json
// @Param input body pb.ActionsRequest true "action info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /account/actions [post]
func (h *Handler) actions(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}

	input := &pb.ActionsRequest{}
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.ID = id
	resp, err := h.authClient.Actions(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{resp.Type})
}
