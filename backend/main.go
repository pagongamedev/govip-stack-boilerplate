package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/pagongamedev/govip-stack-boilerplate/backend/main/middleware"
	"github.com/pagongamedev/govip-stack-boilerplate/backend/main/secret"
	"github.com/pagongamedev/govip-stack-boilerplate/backend/main/storage"

	godd "github.com/pagongamedev/go-dd"
	goddPortal "github.com/pagongamedev/go-dd/portal"
	goddGofiberV2 "github.com/pagongamedev/go-dd/support/gofiber/v2"

	authDomain "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/domain"
	authRepo "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/storage/postgres"
)

func main() {
	portal, secret, stateStorage, _ := goddPortal.New(secret.Get(), storage.Load, middleware.Load)

	appMain := appMain(*stateStorage)

	portal.AppendApp(appMain, ":8081")
	portal.AppendApp(appAPIDocument(*secret), ":8082")
	portal.AppendApp(goddGofiberV2.AppMetricsPrometheus(appMain), ":8083")

	portal.StartServer()
}

func appMain(stateStorage godd.Map) godd.InterfaceApp {
	goddApp, app := goddGofiberV2.NewApp()
	app.Use(cors.New())
	// app.Use(logger.New())

	authDomain.Dock(goddApp, "/income/v1", authRepo.NewRepository, stateStorage["database"])

	return goddApp
}

func appAPIDocument(secret map[string]string) godd.InterfaceApp {

	// docs.SwaggerInfo.Title = secret["project_name"] + " " + secret["doc_service_title"]
	// docs.SwaggerInfo.Version = secret["doc_version"]
	// docs.SwaggerInfo.Description = secret["doc_description"]

	// if secret["doc_api_host"]  != ""{
	// 	docs.SwaggerInfo.host = secret["doc_api_host"]
	// }

	// if secret["doc_basepath"]  != ""{
	// 	docs.SwaggerInfo.basePath = secret["doc_basepath"]
	// }

	// log.Println(secret["doc_service_title"] + " Docs started on :8082")

	return goddGofiberV2.AppAPIDocument()
}
