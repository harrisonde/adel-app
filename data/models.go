package data

import (
	"database/sql"
	"fmt"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	// Models inserted here (and in New fn)
	// are accessible throughout the entire application
	// e.g., Users User
}

// New-up the models for use by the package
func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		upper, _ = postgresql.New(databasePool)
	default:
		// Do nothing as we might no have a database
	}

	return Models{
		// Add your inserted models here, these are the models
		// you add in your models struct
		// e.g., Users:  User{},
	}
}

// Support different DB ID return types
func getInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i) // get type

	// Postgres
	if idType == "int64" {
		return int(i.(int64))
	}

	// Anything else
	return i.(int)
}
