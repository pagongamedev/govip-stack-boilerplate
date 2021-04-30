package auth

import (
	godd "github.com/pagongamedev/go-dd"
)

// Because Loop Cycle Package
type funcRepo = func(database interface{}) (repo Repository, err error)

func Dock(goddApp godd.InterfaceApp, path string, funcRepo funcRepo, database interface{}) {

	repo, err := funcRepo(database)
	godd.MustError(err)
	service, err := NewService(repo)
	godd.MustError(err)

	Router(goddApp, path, service)
}
