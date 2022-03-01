package model

import "fmt"

// todo : add validation func, sanitazie
type Account struct {
	ID           int     `json:"id"`            //1
	UUID         string  `json:"uuid"`          // abc
	WalletAmount float64 `json:"wallet_amount"` //400...
	Currency     int     `json:"currency"`      //rub/usd
}

func (u *User) Validation() error {
	if u == nil {
		return fmt.Errorf("empty User")
	}
	if u.Name == "" {
		return fmt.Errorf("empty name")
	}
	return nil
}
