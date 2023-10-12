package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title      string
	Author     string
	Body       string
	Created_At time.Time
	Likes      int
}
