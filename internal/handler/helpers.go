package handler

import (
	"avito/internal/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func repsonseWithStatus(c *gin.Context, status int, data interface{}, text, message string) {
	c.JSON(status, &model.Response{
		Text:    text,
		Message: message,
		Data:    data,
	})
}

func exchange(fromCurrency, toCurrency string, amount float64) (float64, error) {

	url := fmt.Sprint(urlExchange, "from=", fromCurrency, "&to=", toCurrency, "&amount=", amount, "&api_key=", keyExchange)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	res := model.Currency{}

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return 0, err
	}
	return res.Result.Usd, nil
}
