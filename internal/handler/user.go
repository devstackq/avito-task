package handler

import (
	"avito/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Success", "user created")
		return
	}

	lidAccount, err := h.accountService.NewAccount()
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	user.Account.ID = int(lidAccount)

	lidUser, err := h.userService.CreateUser(&user)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Success", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, lidUser, "Success", "user created")
}
