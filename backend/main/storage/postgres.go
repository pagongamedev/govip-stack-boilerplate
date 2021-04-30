package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	godd "github.com/pagongamedev/go-dd"
)

// LoadPostgres func
func LoadPostgres(secret godd.MapString, storage *godd.Map, deferCloseList *[]godd.DeferClose) (string, error) {

	var commandTemplate = "dbname=%s user=%s password=%s host=%s sslmode=%s application_name=%s"

	if secret["database_sslmode"] == "" {
		log.Println("database_sslmode : disable")
		secret["database_sslmode"] = "disable"
	}

	command := fmt.Sprintf(
		commandTemplate,
		secret["database_name"],
		secret["database_user"],
		secret["database_pass"],
		secret["database_host"],
		secret["database_sslmode"],
		secret["database_name"])

	if secret["database_port"] != "" {
		command = fmt.Sprintf(
			command+" port=%s",
			secret["database_port"])
	}

	if secret["database_sslmode"] == "require" {
		command = fmt.Sprintf(
			command+" sslcert=%s sslkey=%s",
			secret["database_sslcert"],
			secret["database_sslkey"])
	}

	// init db
	db, err := sql.Open(secret["database_driver"], command)
	if err != nil {
		log.Println("Error sql Open:", err)
		return "", err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error Ping:", err)
		return "", err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)
	fmt.Println("connect db success")

	// implement data
	(*storage)["database"] = db
	*deferCloseList = append(*deferCloseList, godd.DeferClose{Name: "SQL", I: db})
	return command, nil
}
