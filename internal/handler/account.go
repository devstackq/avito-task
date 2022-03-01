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
func (h *Handler) AddCurrencyHandler(c *gin.Context) {
}

//user model - check exist; register;
//account - operation
//JSON :uuid, amount, currency; user.uuid: abc, account {wallet:  200, currency: 2(usd)}

func (h *Handler) AddBalance(c *gin.Context) {

	var (
		user model.User
		id   int
		err  error
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

	uuid, err := h.userService.IsExistUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: fmt.Sprint("internal server error: ", err.Error()),
		})
		return
	}
	if uuid != "" {
		c.JSON(http.StatusInternalServerError, &model.Response{
			Text:    "error : ",
			Message: "not found user by uuid, please to register...",
		})
		return
	}

	var balance model.Account

	balance.UUID = user.UUID
	balance.Currency = user.Account.Currency
	balance.WalletAmount = user.Account.WalletAmount

	h.accountService.Add(&balance) //uuid, amount, currnecy; update balance;

	c.JSON(http.StatusOK, &model.Response{
		Text:    "success: ",
		Message: "balance accured",
	})
}

//use transaction
func (h *Handler) DebitBalance(c *gin.Context) {
	var (
		// user model.User
		err error
		id  int
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
	if err != nil {
		log.Print(err)
		c.JSON(500, "error: ...") ///todo
		return
	}
	if uuid != "" {
		//check if currentAmount > debinAmount ?
		h.accountService.Debit()
	} else {
		c.JSON(400, "no have user")
	}
}

func (h *Handler) TransferBalance(c *gin.Context) {
	var senderBalance bool
	//todo validation
	// var user model.User

	uuid, err := h.userService.IsExistUser(0) // senderUuuid check
	if err != nil {
		log.Print(err)
		c.JSON(500, "error: ...") //todo
		return
	}
	//check AmountSum
	if uuid != "" {
		senderBalance, err = h.accountService.CheckBalanceByID() //uuid,
		if err != nil {
			log.Print(err)
			c.JSON(500, "error: ...") //todo
			return
		}
		//if ok -> send else { error; 400; not balance}
	}
	//check reciever

	// uuid, err := h.userService.IsExistUser(0) // receiverUuid
	// if err != nil {
	// 	log.Print(err)
	// 	c.JSON(500, "error: ...") //todo
	// 	return
	// }
	if uuid != "" && senderBalance {
		var balance model.Account

		h.accountService.Add(&balance) // toUUid
		h.accountService.Debit()       // fromUUId
		// h.accountService.Transfer()// fromUUid, toUuid, Amount, updateBlance(0)
	} else {
		c.JSON(400, "no have user")
	}
}
