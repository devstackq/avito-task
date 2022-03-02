package handler

import (
	"avito/internal"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accountService  internal.AccountBalanceServiceInterface
	userService     internal.UserServiceInterface
	currencyService internal.CurrencyServiceInterface
	// cfg             config.Conf
}

func NewHandler(account internal.AccountBalanceServiceInterface, user internal.UserServiceInterface, currency internal.CurrencyServiceInterface) *Handler {
	return &Handler{
		accountService:  account,
		userService:     user,
		currencyService: currency,
	}
}

func SetEnpoints(group *gin.RouterGroup, account internal.AccountBalanceServiceInterface, user internal.UserServiceInterface, currency internal.CurrencyServiceInterface) {

	h := NewHandler(account, user, currency)

	userGroup := group.Group("/user")
	{
		userGroup.POST("/register", h.Register)
		// userGroup.POST("/signin", h.Signin)
	}

	currencyGroup := group.Group("/currency")
	{
		currencyGroup.POST("/", h.CreateCurrency)
	}

	accountGroup := group.Group("/account")
	{
		accountGroup.POST("/add/:id", h.AddBalance)
		accountGroup.POST("/debit/:id", h.DebitBalance)
		accountGroup.POST("/transfer/:id", h.TransferBalance)
		accountGroup.GET("/balance/:id", h.GetBalanceByID)

		accountGroup.POST("/convert/:currency", h.Convert)            //convert rub  to usd
		accountGroup.POST("/currency/:user_id", h.AddCurrencyAccount) //create new currency to user

		// billing.POST("/history/:id", h.GetHistoryTransaction)
	}
}
