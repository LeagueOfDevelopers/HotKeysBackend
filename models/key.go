package models

import (
	"github.com/jinzhu/gorm"
)

type Key struct {
	gorm.Model
	Code int       `json:"code"`
	Name string    `json:"name"`
}
