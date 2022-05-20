package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.client.ParseToken(c, &pb.ParseTokenRequest{Token: headerParts[1]})
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserID(c *gin.Context) (int32, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt32, ok := id.(int32)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return idInt32, nil
}
