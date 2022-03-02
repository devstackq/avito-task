package handler

import (
	"avito/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCurrency(c *gin.Context) {

	var (
		err     error
		account model.Account
	)

	if err = c.ShouldBindJSON(&account); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	err = account.Validation()
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	err = h.currencyService.Create(account.Currency)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, account.Currency, "Success", "create currency")
}
