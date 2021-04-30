package storage

import (
	godd "github.com/pagongamedev/go-dd"
)

// Load func
func Load(secret godd.MapString) (*godd.Map, *[]godd.DeferClose) {
	storage := godd.Map{}
	deferCloseList := []godd.DeferClose{}

	_, err := LoadPostgres(secret, &storage, &deferCloseList)
	godd.MustError(err)

	return &storage, &deferCloseList
}
