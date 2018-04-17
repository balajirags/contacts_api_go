package main

import (
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/contacts_api_go/config"
	"github.com/contacts_api_go/logger"
	"github.com/contacts_api_go/server"
	"github.com/contacts_api_go/appcontext"
	"github.com/urfave/cli"
	"os"
	"github.com/contacts_api_go/console"
)

func main() {
	config.Load()
	logger.InitLogger(config.GetLogLevel())
	appcontext.Initialize()
	clientApp := cli.NewApp()
	clientApp.Name = "contacts"
	clientApp.Version = "0.0.1"
	clientApp.Action = func(c *cli.Context) error {
		server.StartAPIServer()
		return nil
	}
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				err := console.RunDatabaseMigrations()
				if err != nil && err.Error() != "no change" {
					return err
				}
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:        "migrate:run",
			Description: "Running Migration",
			Action: func(c *cli.Context) error {
				err := console.RunDatabaseMigrations()
				if err != nil && err.Error() == "no change" {
					return nil
				} else {
					return err
				}
				return nil
			},
		},
		{
			Name:        "migrate:rollback",
			Description: "Rollback Migration",
			Action: func(c *cli.Context) error {
				console.RollbackLatestMigration()
				return nil
			},
		},
		{
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				return console.CreateMigrationFiles(c.Args().Get(0))
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
