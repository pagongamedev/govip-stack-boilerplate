package auth

import (
	godd "github.com/pagongamedev/go-dd"
)

func Dock(goddApp godd.InterfaceApp, path string, repo Repository) {
	Router(goddApp, path, NewService(repo))
}
