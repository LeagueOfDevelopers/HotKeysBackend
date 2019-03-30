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
	Hotkeys  *[]hotkey.Hotkey `json:"hotkeys"`
}

var ErrorProgramNotFound = errors.New("program not found")
var ErrorGetProgram = errors.New("error while getting program")
var ErrorProgramIdFormat = errors.New("program id wrong format")
var ErrorNoProgramIdFound = errors.New("no param for program id")

type Repository interface {
	GetAll() (*[]Program, error)
	Get(id uint) (*Program, error)
	GetHotkeys(id uint) (*[]hotkey.Hotkey, error)
}
