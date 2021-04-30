package secret

import (
	"os"

	godd "github.com/pagongamedev/go-dd"
)

// GetSentry func
func GetSentry(env godd.Env, secret *godd.MapString) {
	(*secret)["sentry_dsn"] = os.Getenv("SENTRY_DSN")
}
