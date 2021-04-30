package secret

import (
	"os"

	godd "github.com/pagongamedev/go-dd"
)

// GetSwagger func
func GetSwagger(env godd.Env, secret *godd.MapString) {
	(*secret)["doc_service_name"] = os.Getenv("SWAGGER_SERVICE_NAME")
	(*secret)["doc_service_title"] = os.Getenv("SWAGGER_SERVICE_TITLE")
	(*secret)["doc_api_host"] = os.Getenv("SWAGGER_API_HOST")
	(*secret)["doc_basepath"] = os.Getenv("SWAGGER_BASEPATH")
	(*secret)["doc_version"] = os.Getenv("SWAGGER_VERSION")
	(*secret)["doc_description"] = os.Getenv("SWAGGER_DESCRIPTION")
}
