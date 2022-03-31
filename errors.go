package main

import "fmt"

type CreateDatabaseError struct {
	error
	err      error
	database string
}

func (e CreateDatabaseError) Error() string {
	return "failed to create database: " + e.database + "\n" + e.err.Error()
}

type CreateUserError struct {
	error
	err  error
	user string
}

func (e CreateUserError) Error() string {
	return "failed to create user: " + e.user + "\n" + e.err.Error()
}

type GrantPrivilegesTemplateError struct {
	error
	err      error
	database string
	user     string
}

func (e GrantPrivilegesTemplateError) Error() string {
	return "failed to grant privileges on database " + e.database + " on user " + e.user + "\n" + e.err.Error()
}

type ConnectMySQLServerError struct {
	error
	err    error
	dbHost string
	dbPort int
}

func (e ConnectMySQLServerError) Error() string {
	return fmt.Sprintf("failed to connect MySQL Server: root@%s:%d", e.dbHost, e.dbPort) + "\n" + e.err.Error()
}

type CloseMySQLConnecttionError struct {
	err    error
	dbHost string
	dbPort int
}

func (e CloseMySQLConnecttionError) Error() string {
	return fmt.Sprintf("failed to close MySQL Server connection: root@%s:%d", e.dbHost, e.dbPort) + "\n" + e.err.Error()
}

type EmptyEnvironmentVariableError struct {
	error
	env string
}

func (e EmptyEnvironmentVariableError) Error() string {
	return "environment variable " + e.env + " is not set"
}

type InvalidPortError struct {
	error
	portString string
}

func (e InvalidPortError) Error() string {
	return "invalid port " + e.portString
}

type CreateTableFailedError struct {
	error
	err      error
	table    string
	database string
}

func (e CreateTableFailedError) Error() string {
	return "cannot create table " + e.table + " on database " + e.database + "\n" + e.err.Error()
}
