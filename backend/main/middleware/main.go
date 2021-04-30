package middleware

import (
	godd "github.com/pagongamedev/go-dd"
)

// Load func
func Load(secret godd.MapString) (*godd.Map, *[]godd.DeferClose) {
	deferCloseList := []godd.DeferClose{}

	// err := LoadSentry(secret, &deferCloseList)
	// godd.MustError(err)

	return nil, &deferCloseList
}
