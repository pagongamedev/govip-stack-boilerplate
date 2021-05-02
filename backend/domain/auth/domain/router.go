package auth

import (
	"github.com/BurntSushi/toml"
	godd "github.com/pagongamedev/go-dd"
	goddMicroservice "github.com/pagongamedev/go-dd/microservice"
	"golang.org/x/text/language"
)

var ms *goddMicroservice.MicroService

// Router Func
func Router(goddApp godd.InterfaceApp, path string, service interface{}) *goddMicroservice.MicroService {

	i18n := godd.NewI18N(language.English, "toml", toml.Unmarshal,
		godd.Map{
			"en": "domain/income/i18n/active.en.toml",
			"th": "domain/income/i18n/active.th.toml",
		})

	ms = goddMicroservice.New(goddApp, path, service, nil, i18n)

	ms.Post("register", handlerRegisterCreate())
	// ms.Post("transaction/create", handlerIncomeCreate())

	return ms
}
