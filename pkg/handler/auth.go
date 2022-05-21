package handler

import (
	"net/http"

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
// @Failure 400,404 {object} errorResponse
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

	id, err := h.client.CreateUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {
	req := &pb.GenerateTokenRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.client.GenerateToken(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": resp.ID})
}

func (h *Handler) confirmation(c *gin.Context) {

}

func (h *Handler) resend(c *gin.Context) {

}

func (h *Handler) set(c *gin.Context) {

}

func (h *Handler) passwordReset(c *gin.Context) {

}
