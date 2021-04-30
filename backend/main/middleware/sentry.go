package middleware

import (
	"time"

	"github.com/getsentry/sentry-go"
	godd "github.com/pagongamedev/go-dd"
)

func LoadSentry(secret godd.MapString, deferCloseList *[]godd.DeferClose) error {
	if secret["sentry_dsn"] == "" {
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         secret["sentry_dsn"],
		Environment: secret["api_environment"],
		Release:     secret["project_name"] + "@" + secret["api_version"],
	})
	if err != nil {
		return err
		// log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	sentry.CaptureMessage("Start Server : " + secret["api_environment"])

	*deferCloseList = append(*deferCloseList, godd.DeferClose{Name: "Sentry", I: &SentryClose{}})

	return nil
}

type SentryClose struct {
}

func (sc *SentryClose) Close() error {
	sentry.Flush(2 * time.Second)
	return nil
}
