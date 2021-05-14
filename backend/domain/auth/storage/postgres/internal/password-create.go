package auth

import (
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	godd "github.com/pagongamedev/go-dd"
	goddSqlx "github.com/pagongamedev/go-dd/support/database/sqlx"
)

func PasswordCreate(tx *sqlx.Tx, userUUID string, hashsalt string) (string, *godd.Error) {

	query := `
	INSERT INTO 
	  auth.password 
		(user_uuid,hashsalt) 
	VALUES 
		({{Values}})
	RETURNING
		 uuid
		 `

	argList := append([]interface{}{}, userUUID, hashsalt)

	type resStruct struct {
		UUID string `db:"uuid"`
	}
	res, goddErr := goddSqlx.QueryOne(tx, query, argList, &resStruct{})
	if goddErr != nil {

		if goddErr.IsContain("duplicate key value violates unique constraint") {
			// goddErr.IsContainSetError("user_email_uk", "address is exist")
			// goddErr.IsContainSetError("user_phone_uk", "phone is exist")
		}

		return "", goddErr
	}
	if res == nil {
		return "", godd.ErrorNew(http.StatusNotImplemented, errors.New("can't create password"))
	}

	return res.(*resStruct).UUID, nil
}
