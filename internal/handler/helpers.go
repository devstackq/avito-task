package handler

import (
	"avito/internal/model"

	"github.com/gin-gonic/gin"
)

func setCredsAccount(account *model.Account, uuid, currency string) {
	account.UUID = uuid
	// account.WalletAmount = amount
	account.Currency = currency
}

func repsonseWithStatus(c *gin.Context, status int, data interface{}, text, message string) {
	c.JSON(status, &model.Response{
		Text:    text,
		Message: message,
		Data:    data,
	})
}
