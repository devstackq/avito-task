package handler

import (
	"avito/internal/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//in furher..; uuid, typeCurrency; for moderator
// func (h *Handler) AddCurrencyHandler(c *gin.Context) {
// }

//user model - check exist; register;
//account - operation
//JSON :uuid, amount, currency; user.uuid: abc, account {wallet:  200, currency: 2(usd)}

//todo: validation, responseWithStatus

const defaultCurrency = "rub"

func (h *Handler) Convert(c *gin.Context) {
	//id user; name currency;
	var (
		currency string
		err      error
		user     model.User
		uuid     string
		acc      model.Account
	)

	currency = c.Param("currency")

	if currency == "" {
		//todo: mini func
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("can't atoi : ", err.Error()),
		})
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "Error",
			Message: fmt.Sprint("bad request: user: ", err.Error()),
		})
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
	//inner service ?
	acc.Currency = currency // to usd
	acc.UUID = uuid

	//todo: logic - in service
	// h.accountService.Convert()

	temp, err := h.accountService.CheckBalanceByUUID(uuid)

	//usd = api(walletAmount)
	acc.WalletAmount = 92.332 //newUsdVal

	//all balance convert
	err = h.accountService.Add(&acc)
	//all rub debit
	acc.WalletAmount = temp
	acc.Currency = defaultCurrency

	err = h.accountService.Debit(&acc)
	//debit(rub, typeCurrency) - rub
	//add(usd, typeCurrency)
	//uuid - 2 more currency
}
func (h *Handler) AddBalance(c *gin.Context) {

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
	// balance.Currency = user.Account.Currency
	balance.WalletAmount = user.Account.WalletAmount
	balance.Currency = defaultCurrency

	h.accountService.Add(&balance) //uuid, amount, currnecy; update balance;

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: "balance accured",
	})
}

fix transfer;
refactor another handlers like Transfer
start convert

//use transaction
//refactor like Transfer
func (h *Handler) DebitBalance(c *gin.Context) {
	var (
		user          model.User
		err           error
		id            int
		balanceAmount float64
		balance       model.Account
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("can't atoi : ", err.Error()),
		})
		return
	}

	uuid, err := h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		log.Print(err)
		c.JSON(500, "error: ...")
		return
	}
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("bad request: user: ", err.Error()),
		})
		return
	}
	//check if currentAmount > debinAmount ?
	balanceAmount, err = h.accountService.CheckBalanceByUUID(uuid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "Error",
			Message: fmt.Sprint("internal server error: ", err.Error()),
		})
		return
	}

	if balanceAmount-user.Account.WalletAmount > 0 {
		balance.Currency = defaultCurrency

		balance.UUID = uuid
		balance.Currency = user.Account.Currency
		balance.WalletAmount = user.Account.WalletAmount

		if err = h.accountService.Debit(&balance); err != nil {
			c.JSON(http.StatusInternalServerError, &model.Response{
				Text:    "Error",
				Message: fmt.Sprint("internal server error: ", err.Error()),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "Error",
			Message: fmt.Sprint("amount nedostaochno: "),
		})
		return
	}

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: "balance recorded",
	})
}

//helper func ?
func (h *Handler) TransferBalance(c *gin.Context) {

	var (
		uuidSender      string
		uuidReceiver    string
		err             error
		id              int
		user            model.User
		balanceSender   model.Account
		balanceReceiver model.Account
	)
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
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

	uuidSender, err = h.userService.IsExistUser(id) // senderUuuid check
	if err != nil {
		log.Print(err)
		c.JSON(500, err.Error())
		return
	}

	//func ()?
	balanceSender.UUID = uuidSender
	balanceSender.Currency = user.Account.Currency
	balanceSender.TransferAmount = user.Account.WalletAmount // debit - 10
	// balanceSender.Currency = "rub"
	balanceSender.Currency = defaultCurrency

	uuidReceiver, err = h.userService.IsExistUser(user.Account.ReceiverID) // senderUuuid check
	if err != nil {
		log.Print(err)
		c.JSON(500, "error: ...")
		return
	}

	balanceReceiver.UUID = uuidReceiver
	balanceReceiver.Currency = user.Account.Currency
	balanceReceiver.TransferAmount = user.Account.WalletAmount //add 10
	// balanceReceiver.Currency = "rub"

	// balanceSender.WalletAmount = user.Account.WalletAmount

	h.accountService.Transfer(balanceSender, balanceReceiver)

	// balanceAmountSender, err = h.accountService.CheckBalanceByUUID(uuidSender)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, &model.Response{
	// 		Text:    "Error",
	// 		Message: fmt.Sprint("internal server error: ", err.Error()),
	// 	})
	// 	return
	// }

	// if balanceAmountSender-user.Account.WalletAmount > 0 {

	// 	balanceSender.UUID = uuidSender
	// 	balanceSender.Currency = user.Account.Currency
	// 	balanceSender.WalletAmount = user.Account.WalletAmount

	// 	if err = h.accountService.Debit(&balanceSender); err != nil {
	// 		c.JSON(http.StatusInternalServerError, &model.Response{
	// 			Text:    "Error",
	// 			Message: fmt.Sprint("internal server error: ", err.Error()),
	// 		})
	// 		return
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, &model.Response{
	// 		Text:    "Error",
	// 		Message: fmt.Sprint("amount nedostaochno: "),
	// 	})
	// 	return
	// }

	// // balanceAmountReceiver, err = h.accountService.CheckBalanceByUUID(uuidReceiver)

	// balanceReceiver.UUID = uuidReceiver
	// balanceReceiver.Currency = user.Account.Currency
	// balanceReceiver.WalletAmount = user.Account.WalletAmount

	// if err = h.accountService.Add(&balanceReceiver); err != nil {
	// 	c.JSON(http.StatusInternalServerError, &model.Response{
	// 		Text:    "Error",
	// 		Message: fmt.Sprint("internal server error: ", err.Error()),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, &model.Response{
		Text:    "Success ",
		Message: "balance transfer",
	})

}

//todo": get blaance id AND currencyType

func (h *Handler) GetBalanceByID(c *gin.Context) {

	var (
		id     int
		err    error
		uuid   string
		amount float64
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("can't atoi : ", err.Error()),
		})
		return
	}
	uuid, err = h.userService.IsExistUser(id)
	if err != nil || uuid == "" {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("internal server error: ", err.Error()),
		})
		return
	}

	amount, err = h.accountService.CheckBalanceByUUID(uuid) //uuid, amount, currnecy; update balance;

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: "get balance",
		Data:    amount,
	})
}
