package service

import "avito/internal"

type CurrencyService struct {
	repo internal.CurrencyRepositoryInterface
}

func (cr CurrencyService) Create(name string) error {
	return cr.repo.Create(name)
}
func (cr CurrencyService) GetCurrencyID(name string) (int, error) {
	return cr.repo.GetCurrencyID(name)
}

func NewCurrencyService(repo internal.CurrencyRepositoryInterface) internal.CurrencyServiceInterface {
	return &CurrencyService{repo: repo}
}
