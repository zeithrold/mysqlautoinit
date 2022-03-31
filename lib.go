package main

import (
	"database/sql"
	"fmt"
	"github.com/dchest/uniuri"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

const MySQLCreateDatabaseTemplate = "CREATE DATABASE %s"
const MySQLCreateUserTemplate = "CREATE USER %s@'%%' IDENTIFIED BY '%s'"
const MySQLGrantPrivilegesTemplate = "GRANT ALL PRIVILEGES ON %s.* TO %s@'%%'"

func CreateDatabases(dbHost string, dbPort int, rootPassword string, databases []string) (map[string]string, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/", rootPassword, dbHost, strconv.Itoa(dbPort)))
	err = db.Ping()
	if err != nil {
		return nil, ConnectMySQLServerError{
			err:    err,
			dbHost: dbHost,
			dbPort: dbPort,
		}
	}
	defer db.Close()
	return CreateDatabasesWithDB(db, databases)
}

func CreateDatabasesWithDB(db *sql.DB, databases []string) (map[string]string, error) {
	result := map[string]string{}
	for _, database := range databases {
		_, err := db.Query(fmt.Sprintf(MySQLCreateDatabaseTemplate, database))
		if err != nil {
			return nil, CreateDatabaseError{database: database, err: err}
		}
		randomPassword := uniuri.NewLen(16)
		_, err = db.Query(fmt.Sprintf(MySQLCreateUserTemplate, database, randomPassword))
		if err != nil {
			return nil, CreateUserError{user: database, err: err}
		}
		_, err = db.Query(fmt.Sprintf(MySQLGrantPrivilegesTemplate, database, database))
		if err != nil {
			return nil, GrantPrivilegesTemplateError{err: err, database: database, user: database}
		}
		result[database] = randomPassword
	}
	return result, nil
}
