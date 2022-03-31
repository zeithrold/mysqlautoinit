package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Usage: "An automation tool to create MySQL databases(and users) rapidly",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "db-host",
				Aliases:     []string{"H"},
				Usage:       "The MySQL Server Host that the tool can access",
				DefaultText: "localhost",
				Required:    true,
			},
			&cli.IntFlag{
				Name:        "db-port",
				Aliases:     []string{"P"},
				Usage:       "The MySQL Server Port that the tool can access",
				DefaultText: "3306",
				Required:    true,
			},
			&cli.StringFlag{
				Name:    "root-password",
				Aliases: []string{"R"},
				Usage:   "The MySQL Server Root Password",
			},
			&cli.StringSliceFlag{
				Name:        "databases",
				Aliases:     []string{"D"},
				Usage:       "The List of the databases you want to create",
				DefaultText: "-D wordpress -D ghost",
				Required:    true,
			},
		},
		Action: func(context *cli.Context) error {
			dbHost := context.String("db-host")
			dbPort := context.Int("db-port")
			rootPassword := context.String("root-password")
			databases := context.StringSlice("databases")
			result, err := CreateDatabases(dbHost, dbPort, rootPassword, databases)
			if err != nil {
				return err
			}
			marshaled, err := json.Marshal(result)
			fmt.Println(string(marshaled))
			return nil
		},
	}
	app.Run(os.Args)
}
