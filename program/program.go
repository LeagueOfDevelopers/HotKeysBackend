package program

import (
	"HotKeysBackend/hotkey"
	"errors"
	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	Name     string           `json:"name"`
	ImageURL string           `json:"url"`
	Hotkeys  []*hotkey.Hotkey `json:"hotkeys"`
}

var ErrorProgramNotFound = errors.New("program not found")

type Repository interface {
	GetAll() ([]*Program, error)
	Get(string) (*Program, error)
}
