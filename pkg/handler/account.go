package handler

import (
	"fmt"
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
// @Param input body pb.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /account/update [post]
func (h *Handler) update(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}

	input := &pb.User{}
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Email == "" && input.Expires == "" && input.Password == "" && input.Type == "" && input.Username == "" {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.ID = id
	fmt.Printf("%s %s", input.Username, input.Password)

	_, err = h.client.Update(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

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
