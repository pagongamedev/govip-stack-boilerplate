package auth

import (
	godd "github.com/pagongamedev/go-dd"
	auth "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/domain"
	authRepo "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/storage/postgres"
)

func Dock(goddApp godd.InterfaceApp, path string, database interface{}) {

	repo := authRepo.NewRepository(database)
	service := auth.NewService(repo)
	auth.Router(goddApp, path, service)
}
