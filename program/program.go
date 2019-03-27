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
var ErrorGetProgram = errors.New("error while getting program")

type Repository interface {
	GetAll() ([]*Program, error)
	Get(name string) (*Program, error)
	GetHotkeys(name string) ([]*hotkey.Hotkey, error)
}
