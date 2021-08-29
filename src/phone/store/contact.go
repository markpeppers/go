package store

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Number string `gorm:"check:unique"`
}
