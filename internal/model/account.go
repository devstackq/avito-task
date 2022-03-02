package model

import "fmt"

type Account struct {
	ID           int     `json:"id"`
	UUID         string  `json:"uuid"`
	WalletAmount float64 `json:"wallet_amount"`
	Currency     string  `json:"currency_name"`
	ReceiverID   int     `json:"receiver_id"`
	CurrencyType int     `json:"currency_type"`
}

func (a *Account) Validation() error {
	if a == nil {
		return fmt.Errorf("empty Account")
	}
	if a.UUID == "" {
		return fmt.Errorf("empty uuid")
	}
	if a.WalletAmount == 0 {
		return fmt.Errorf("empty amount")
	}
	if a.Currency == "" {
		return fmt.Errorf("empty currency")
	}
	return nil
}
