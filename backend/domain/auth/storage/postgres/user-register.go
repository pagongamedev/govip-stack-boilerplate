package auth

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
	internal "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/storage/postgres/internal"
)

func (r *repo) UserRegister(username string, hashsalt string, phone string) (interface{}, *godd.Error) {
	tx := r.db.MustBegin()
	// response := godd.Map{"User": username, "Password": hashsalt, "Phone": phone}

	userUUID, goddErr := internal.UserCreate(tx, username, phone)
	if goddErr != nil {
		return nil, goddErr
	}

	_, goddErr = internal.PasswordCreate(tx, userUUID, hashsalt)
	if goddErr != nil {
		return nil, goddErr
	}

	err := tx.Commit()
	if err != nil {
		return nil, godd.ErrorNew(http.StatusNotImplemented, err)

	}

	return userUUID, nil
}
