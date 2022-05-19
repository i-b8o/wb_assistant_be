package handler

import (
	"fmt"
	"net/http"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("AAAAAAAAAAAAAAAA")
	input := &pb.User{}
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

}

func (h *Handler) confirmation(c *gin.Context) {

}

func (h *Handler) resend(c *gin.Context) {

}

func (h *Handler) set(c *gin.Context) {

}

func (h *Handler) passwordReset(c *gin.Context) {

}
