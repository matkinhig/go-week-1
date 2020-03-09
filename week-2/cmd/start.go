package cmd

import (
	"github.com/Team/go-week-2/myservice"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
)

// 1. Load du lieu url tu trong db ra
// 2. Lay cai url crawl du lieu
// 3. Parse cai moi dung html crawl dc ve thanh cau truc
// title. content, published date, author (Article)
// 4. Insert du lieu vao val trong db

// Start comment
var Start = cli.Command{
	Name:  "start",
	Usage: "Start application",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		svc := myservice.NewCrawlerService(db)
		// Doc lap chi load urls tu trong db
		urls := svc.Load()
		// Neu ma co urls tu trong db => no se crawl
		articles := svc.Crawl(urls)
		svc.InsertArcitlesToDB(articles)
		return nil
	},
}
