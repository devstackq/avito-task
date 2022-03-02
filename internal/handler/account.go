package handler

import (
	"avito/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultCurrency = "rub"

//todo convert; unit test; docker

func (h *Handler) Convert(c *gin.Context) {
	var (
		currency string
		err      error
		user     model.User
		uuid     string
		acc      model.Account
	)

	if currency = c.Param("currency"); currency == "" {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", "empty currency")
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	uuid, err = h.userService.IsExistUser(user.ID)
	if err != nil || uuid == "" {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("internal server error: ", err.Error()),
		})
		return
	}
	acc.Currency = currency // to usd
	acc.UUID = uuid

	//todo: logic - in service
	// h.accountService.Convert()
	// temp, err := h.accountService.CheckBalance(uuid)
	//usd = api(walletAmount)
	acc.WalletAmount = 92.332 //newUsdVal
	//all balance convert
	err = h.accountService.Add(&acc)
	//all rub debit
	// acc.WalletAmount = temp
	acc.Currency = defaultCurrency
	err = h.accountService.Debit(&acc)
	//debit(rub, typeCurrency) - rub
	//add(usd, typeCurrency)
	//uuid - 2 more currency
}

func (h *Handler) AddBalance(c *gin.Context) {
	var (
		id      int
		err     error
		balance model.Account
		uuid    string
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	if err = c.ShouldBindJSON(&balance); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	uuid, err = h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	balance.UUID = uuid
	balance.Currency = defaultCurrency

	err = balance.Validation()

	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	err = h.accountService.Add(&balance)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, balance.WalletAmount, "Success", "Balance occured")
}

func (h *Handler) DebitBalance(c *gin.Context) {
	var (
		err     error
		id      int
		balance model.Account
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	uuid, err := h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	if err = c.ShouldBindJSON(&balance); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	balance.Currency = defaultCurrency
	balance.UUID = uuid

	if err = h.accountService.Debit(&balance); err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, balance.WalletAmount, "Success", "Balance recorded")
}

func (h *Handler) TransferBalance(c *gin.Context) {

	var (
		uuidSender      string
		uuidReceiver    string
		err             error
		id              int
		balanceSender   model.Account
		balanceReceiver model.Account
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	uuidSender, err = h.userService.IsExistUser(id)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	if err = c.ShouldBindJSON(&balanceSender); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	balanceSender.UUID = uuidSender
	balanceSender.Currency = defaultCurrency

	uuidReceiver, err = h.userService.IsExistUser(balanceSender.ReceiverID) // senderUuuid check
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	balanceReceiver.UUID = uuidReceiver
	balanceReceiver.Currency = balanceSender.Currency
	balanceReceiver.WalletAmount = balanceSender.WalletAmount

	err = h.accountService.Transfer(balanceSender, balanceReceiver)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, balanceSender.WalletAmount, "Success", "Amount transfered")
}

func (h *Handler) GetBalanceByID(c *gin.Context) {

	var (
		id      int
		err     error
		uuid    string
		amount  float64
		account model.Account
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	uuid, err = h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	if err = c.ShouldBindJSON(&account); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	amount, err = h.accountService.CheckBalance(uuid, account.CurrencyType) //uuid, amount, currnecy; update balance;
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, amount, "Success", "Get balance")

}
