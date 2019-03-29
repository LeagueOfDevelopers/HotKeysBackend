package program

import (
	"HotKeysBackend/hotkey"
	"errors"
)

type Program struct {
	Id       uint             `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name     string           `json:"name"`
	ImageURL string           `json:"url"`
	Hotkeys  *[]hotkey.Hotkey `json:"hotkeys"`
}

var ErrorProgramNotFound = errors.New("program not found")
var ErrorGetProgram = errors.New("error while getting program")
var ErrorProgramIdFormat = errors.New("program id wrong format")

type Repository interface {
	GetAll() (*[]Program, error)
	Get(id uint) (*Program, error)
	GetHotkeys(id uint) (*[]hotkey.Hotkey, error)
}
