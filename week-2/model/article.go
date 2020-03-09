package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article struct
type Article struct {
	gorm.Model
	URLID       uint
	Title       string
	Content     string
	PublishedAt time.Time
	Author      string
}
