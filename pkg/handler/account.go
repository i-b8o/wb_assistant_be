package handler

import (
	"fmt"
	"net/http"

	"github.com/bogach-ivan/nonsense"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

// @Summary update email verification token
// @Security ApiKeyAuth
// @Tags account
// @Description update email verification token
// @ID update-email-verification-token
// @Accept json
// @Produce json
// @Param input body pb.UpdateEmailVerificationTokenRequest true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /account/update-email-verification-token [post]
func (h *Handler) updateEmailVerificationToken(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		return
	}

	input := &pb.UpdateEmailVerificationTokenRequest{}
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token := nonsense.RandSeq(100)
	input.Token = token
	input.ID = id

	_, err = h.authClient.UpdateEmailVerificationToken(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Send the confirm token to a user email
	fmt.Println("bdrop.net/auth/confirmation/" + token + "e:" + input.Email + "P:" + input.Password)
	r, err := h.mailClient.Confirm(c, &pb.MailConfirmRequest{Url: "bdrop.net/auth/confirmation/" + token, Email: input.Email, Pass: input.Password})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(r.Message) > 0 {
		newErrorResponse(c, http.StatusInternalServerError, r.Message)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

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
