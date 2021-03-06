package handler

import (
	"net/http"
	"net/mail"

	"github.com/bogach-ivan/nonsense"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body pb.CreateUserRequest true "account info"
// @Success 200 {object} pb.CreateUserResponse 1
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	input := &pb.CreateUserRequest{}
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = mail.ParseAddress(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Create User
	// TODO separate user exist from error
	resp, err := h.authClient.CreateUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if resp.ID == 0 {
		newErrorResponse(c, http.StatusConflict, "")
	}

	// Add confirm token to db
	token := nonsense.RandSeq(100)
	_, err = h.authClient.InsertEmailConfirmToken(c, &pb.InsertEmailConfirmTokenRequest{ID: resp.ID, Token: token})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Send the confirm token to a user email
	r, err := h.mailClient.Confirm(c, &pb.MailConfirmRequest{Url: "bdrop.net/auth/confirmation/" + token, Email: input.Email, Pass: input.Password})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(r.Message) > 0 {
		newErrorResponse(c, http.StatusInternalServerError, r.Message)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": resp.ID})
}

// @Summary SignIn
// @Tags auth
// @Description sign in to account
// @ID sign-in-account
// @Accept json
// @Produce json
// @Param input body pb.GenerateTokenRequest true "account info"
// @Success 200 {object} pb.GenerateTokenResponse 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	req := &pb.GenerateTokenRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Email validation
	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Try to get jwt
	resp, err := h.authClient.GenerateToken(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(resp.Token) == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": resp.Token})
}

// @Summary Password recover
// @Tags recover
// @Description recover password
// @ID recover
// @Accept json
// @Produce json
// @Param input body pb.RecoverPasswordRequest true "account info"
// @Success 200 {object} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/recover [post]
func (h *Handler) passwordRecover(c *gin.Context) {
	req := &pb.RecoverPasswordRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Email validation
	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Generete new password
	pass := nonsense.RandSeq(7)
	req.Password = pass
	_, err = h.authClient.RecoverPassword(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Send the password to the user email
	r, err := h.mailClient.Reset(c, &pb.ResetRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	if len(r.Message) > 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.Writer.WriteHeader(200)
}

func (h *Handler) confirmation(c *gin.Context) {

	token := c.Param("token")

	_, err := h.authClient.CheckAndDelEmailConfirmToken(c, &pb.CheckAndDelEmailConfirmTokenRequest{
		Token: token,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusFound, "https://www.google.com/")
	c.Abort()
}

// @Summary update email verification token
// @Tags auth
// @Description update email verification token
// @ID update-email-verification-token
// @Accept json
// @Produce json
// @Param input body pb.UpdateEmailVerificationTokenRequest true "auth info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/update-email-verification-token [post]
func (h *Handler) updateEmailVerificationToken(c *gin.Context) {

	input := &pb.UpdateEmailVerificationTokenRequest{}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token := nonsense.RandSeq(100)
	input.Token = token

	_, err := h.authClient.UpdateEmailVerificationToken(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

func (h *Handler) resend(c *gin.Context) {

}

func (h *Handler) set(c *gin.Context) {

}
