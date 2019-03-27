package program

import (
	"HotKeysBackend/hotkey"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Program struct {
	gorm.Model
	Id       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	ImageURL string           `json:"url"`
	Hotkeys  []*hotkey.Hotkey `json:"hotkeys"`
}

var ErrorProgramNotFound = errors.New("program not found")

type Repository interface {
	GetAll() ([]*Program, error)
	Get(string) (*Program, error)
}
