package auth

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
)

func (r *repo) UserRegister(username string, hashsalt string, phone string) (interface{}, *godd.Error) {
	tx := r.db.MustBegin()
	response := godd.Map{"User": username, "Password": hashsalt, "Phone": phone}
	// bookType := "USER"

	// if isExchange {
	// 	bookType = "EXCH"
	// }

	// accBalance, goddErr := r.InternalCreate(tx, exchangeID, currencyCode, bookType, address, name)
	// if goddErr != nil {
	// 	return nil, goddErr
	// }

	err := tx.Commit()
	if err != nil {
		return nil, godd.ErrorNew(http.StatusNotImplemented, err)

	}

	// if accBalance == nil {
	// 	return nil, godd.ErrorNew(http.StatusNotFound, errors.New("address not exist"))
	// }

	// response := &wallet.MappingBookBalance{
	// 	IsLock:       accBalance.IsLock,
	// 	BookType:     accBalance.BookType,
	// 	ExchangeName: accBalance.ExchangeName,
	// 	AddressName:  accBalance.AddressName,
	// 	CurrencyCode: accBalance.CurrencyCode,
	// 	Address:      accBalance.Address,
	// 	Balance:      r.HelperIntToFloat(accBalance.Balance),
	// 	CreatedAt:    accBalance.CreatedAt,
	// }

	return response, nil
}
