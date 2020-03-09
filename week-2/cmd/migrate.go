package cmd

import (
	"github.com/Team/go-week-2/model"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
)

var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrate schema to database",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		db.DropTableIfExists(&model.Url{}, &model.Article{})
		db.AutoMigrate(&model.Url{}, &model.Article{})
		return nil
	},
}
