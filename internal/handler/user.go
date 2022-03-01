package handler

import (
	"avito/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//todo : func helper
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("bad request: user: ", err.Error()),
		})
		return
	}

	lidAccount, err := h.accountService.NewAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("can't create new account: ", err.Error()),
		})
		return
	}

	user.Account.ID = int(lidAccount)

	lidUser, err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text: "E	rror",
			Message: fmt.Sprint("can't create new user: ", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: fmt.Sprint("create new user by id: ", lidUser),
	})
}
