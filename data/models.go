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

type Models struct{
	/*
	|--------------------------------------------------------------------------
	| Models
	|--------------------------------------------------------------------------
	|
	| Here is where you can insert you models for the application. These
	| models are accessible throughout the entire application. For example
	| you may insert "Users User" here. You will need to return your newly
	| created models in the return "Users: User{}" function too!
	|
	*/

	
}

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

	return Models{}
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
