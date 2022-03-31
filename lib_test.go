package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"testing"
)

var db *sql.DB
var err error

const DefaultMySQLHost = "localhost"
const DefaultMySQLPort = "3306"
const DefaultRootPassword = ""

var TestDatabases = []string{
	"GOODTEK",
	"REIMEI",
	"BetterGraphicAnimation",
}

var globalDBHost string
var globalDBPort int

func TestMain(m *testing.M) {
	dbHost := os.Getenv("Z_TEST_MYSQL_HOST")
	dbPortString := os.Getenv("Z_TEST_MYSQL_PORT")
	rootPassword := os.Getenv("Z_TEST_ROOT_PASSWORD")
	if dbHost == "" {
		dbHost = DefaultMySQLHost
	}
	if dbPortString == "" {
		dbPortString = DefaultMySQLPort
	}
	if rootPassword == "" {
		rootPassword = DefaultRootPassword
	}
	var dbPort int
	dbPort, err = strconv.Atoi(dbPortString)
	if err != nil {
		handleError(InvalidPortError{portString: dbPortString})
	}
	globalDBHost = dbHost
	globalDBPort = dbPort
	db, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%d)/", rootPassword, dbHost, dbPort))
	if err != nil {
		handleError(ConnectMySQLServerError{
			err:    err,
			dbHost: dbHost,
			dbPort: dbPort,
		})
	}
	m.Run()
	err = db.Close()
	if err != nil {
		handleError(CloseMySQLConnecttionError{
			err:    err,
			dbHost: dbHost,
			dbPort: dbPort,
		})
	}
}

func TestCreateDatabases(t *testing.T) {
	var result map[string]string
	result, err = CreateDatabasesWithDB(db, TestDatabases)
	if err != nil {
		t.Fatal(err)
	}
	for database, password := range result {
		var anotherDB *sql.DB
		anotherDB, err = sql.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", database, password, globalDBHost, globalDBPort, database))
		err = anotherDB.Ping()
		if err != nil {
			t.Fatal(ConnectMySQLServerError{err: err, dbHost: globalDBHost, dbPort: globalDBPort})
		}
		_, err = anotherDB.Query(fmt.Sprintf("CREATE TABLE %s (id INT PRIMARY KEY, name VARCHAR(255))", database))
		if err != nil {
			t.Fatal(CreateTableFailedError{
				err:      err,
				table:    database,
				database: database,
			})
		}
	}
}

func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(-1)
}
