package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Hotkey struct {
	gorm.Model
	Id          uuid.UUID `json:"id"`
	Description string    `json:"desc"`
	Combination []Key     `json:"combination"`
}
