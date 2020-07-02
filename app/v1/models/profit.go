package models

import "github.com/jinzhu/gorm"

type Profit struct {
	gorm.Model
	Profit string
	Uid string
}
