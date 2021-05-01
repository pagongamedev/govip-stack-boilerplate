package income

import (
	godd "github.com/pagongamedev/go-dd"
	income "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/income/domain"
	incomeRepo "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/income/storage/postgres"
)

func Dock(goddApp godd.InterfaceApp, path string, database interface{}) {

	repo := incomeRepo.NewRepository(database)
	service := income.NewService(repo)
	income.Router(goddApp, path, service)
}
