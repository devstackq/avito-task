package handler

import (
	"avito/internal"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accountService internal.AccountBalanceServiceInterface
	userService    internal.UserServiceInterface
}

func NewHandler(account internal.AccountBalanceServiceInterface, user internal.UserServiceInterface) *Handler {
	return &Handler{
		accountService: account,
		userService:    user,
	}
}

func SetEnpoints(group *gin.RouterGroup, account internal.AccountBalanceServiceInterface, user internal.UserServiceInterface) {

	h := NewHandler(account, user)

	userGroup := group.Group("/user")
	{
		userGroup.POST("/register", h.Register)
		// userGroup.POST("/signin", h.Signin)
	}

	currencyGroup := group.Group("/currency")
	{
		currencyGroup.POST("/", h.CreateCurrency)
	}

	billingGroup := group.Group("/billing")
	{
		billingGroup.POST("/add/:id", h.AddBalance)
		billingGroup.POST("/debit/:id", h.DebitBalance)
		billingGroup.POST("/transfer/:id", h.TransferBalance)
		billingGroup.GET("/balance/:id", h.GetBalanceByID)

		billingGroup.POST("/:currency", h.Convert) //convert

		//todo:
		// billing.POST("/history/:id", h.GetHistoryTransaction)
	}
}
