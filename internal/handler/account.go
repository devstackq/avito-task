package handler

import (
	"avito/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultCurrency = "rub"
	urlExchange     = "https://api.fastforex.io/convert?"
	keyExchange     = "cb0160a3ce-2050986345-r846ss"
)

func (h *Handler) AddCurrencyAccount(c *gin.Context) {

	var (
		err     error
		id      int
		account model.Account
		uuid    string
	)

	id, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	if err = c.ShouldBindJSON(&account); err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	uuid, err = h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	err = h.accountService.AddCurrencyAccount(uuid, account.Currency)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, account.Currency, "Success", "added new currency ")

}

func (h *Handler) Convert(c *gin.Context) {
	var (
		currency string
		err      error
		uuid     string
		user     model.User
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
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	acc.Currency = defaultCurrency // to usd
	acc.UUID = uuid
	acc.ConvertCurrency = currency
	//get current amount - rub
	acc.CurrencyType, err = h.currencyService.GetCurrencyID(acc.Currency)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	amount, err := h.accountService.CheckBalance(uuid, acc.CurrencyType)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	if amount <= 0 {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", "nedostatochno sredstv dlya perevoda")
		return
	}

	converted, err := exchange(acc.Currency, acc.ConvertCurrency, amount)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}

	acc.ConvertedAmount = converted

	err = h.accountService.Convert(&acc)
	if err != nil {
		repsonseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, converted, "Success", "converted ")
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

	amount, err = h.accountService.CheckBalance(uuid, account.CurrencyType)
	if err != nil {
		repsonseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	repsonseWithStatus(c, http.StatusOK, amount, "Success", "Get balance")

}
