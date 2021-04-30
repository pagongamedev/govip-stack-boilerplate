package secret

import (
	"os"

	godd "github.com/pagongamedev/go-dd"
)

// GetPostgres func
func GetPostgres(env godd.Env, secret *godd.MapString) {
	(*secret)["database_driver"] = "postgres"

	(*secret)["database_name"] = os.Getenv("DB_NAME")
	(*secret)["database_host"] = os.Getenv("DB_HOST")
	(*secret)["database_user"] = os.Getenv("DB_USER")
	(*secret)["database_pass"] = os.Getenv("DB_PASS")
	(*secret)["database_port"] = os.Getenv("DB_PORT")
	(*secret)["database_sslmode"] = os.Getenv("DB_SSLMODE")
	(*secret)["database_sslcert"] = os.Getenv("DB_SSLCERT")
	(*secret)["database_sslkey"] = os.Getenv("DB_SSLKEY")

	// file, err := os.Open((*secret)["database_sslkey"])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err = file.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// b, err := ioutil.ReadAll(file)
	// fmt.Printf("\n\nKey\n\n")
	// fmt.Println(string(b))

	// file2, err := os.Open((*secret)["database_sslcert"])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err = file2.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// b, err = ioutil.ReadAll(file2)
	// fmt.Printf("\n\nCert\n\n")
	// fmt.Println(string(b))

	switch env {
	case godd.EnvLocalhost:

		break
	case godd.EnvDevelopment:

		break
	case godd.EnvTesting:

		break
	case godd.EnvStaging:

		break
	case godd.EnvProduction:

		break
	}

}
