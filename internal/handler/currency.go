package handler

import (
	"avito/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCurrency(c *gin.Context) {

	var (
		user    model.User
		id      int
		err     error
		balance model.Account
		uuid    string
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		//todo: mini func
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("can't atoi : ", err.Error()),
		})
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("bad request: user: ", err.Error()),
		})
		return
	}

	//todo: model. validation, sanitaze()

	uuid, err = h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("internal server error: ", err.Error()),
		})
		return
	}

	balance.UUID = uuid
	balance.Currency = user.Account.Currency
	balance.WalletAmount = user.Account.WalletAmount

	h.accountService.Add(&balance) //uuid, amount, currnecy; update balance;

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: "balance accured",
	})
}
