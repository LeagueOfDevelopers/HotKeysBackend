package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Program struct {
	gorm.Model
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ImageURL string    `json:"url"`
	Hotkeys  []Hotkey  `json:"hotkeys"`
}
