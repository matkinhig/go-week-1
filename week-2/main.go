package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Team/go-week-2/cmd"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &cli.App{
		Name:    "spider",
		Usage:   "Crawler Service",
		Version: "1.0.0",
		Action: func(c *cli.Context) error {
			fmt.Println("Welcome to Crawler Service. To see how to use, type -h for more detail")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db", "d"},
				EnvVars: []string{"DB_URI"},
			},
			&cli.BoolFlag{
				Name:    "dblog",
				Aliases: []string{"db_log"},
				EnvVars: []string{"DB_LOG"},
				Value:   false,
			},
		},
		Before: func(c *cli.Context) error {
			//database = "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
			db, err := gorm.Open("mysql", c.String("database"))
			if err != nil {
				panic(err)
			}
			db.LogMode(c.Bool("dblog"))
			c.App.Metadata["db"] = db
			return nil
		},
		Commands: []*cli.Command{&cmd.Migrate, &cmd.Start, &cmd.Seed},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
