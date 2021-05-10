package secret

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	godd "github.com/pagongamedev/go-dd"
)

// Get func
func Get() *godd.MapString {
	secret := godd.MapString{}

	env := os.Getenv("API_ENV")
	version := os.Getenv("API_VERSION")

	// ======= Load Dot ENV =======
	log.Println("App Mode : " + env)

	err := godotenv.Load()
	if err == nil {
		log.Printf("Load .ENV\n")
	} else {
		log.Printf("Not Load .ENV\n")
	}

	// ============================
	secret["api_environment"] = env

	if version == "" {
		version = "0.0.0"
	}
	secret["api_version"] = version

	secret["project_name"] = os.Getenv("PROJECT_NAME")

	GetPostgres(godd.Env(env), &secret)
	GetSwagger(godd.Env(env), &secret)
	GetSentry(godd.Env(env), &secret)

	return &secret
}
