package program

import (
	"HotKeysBackend/hotkey"
	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	Name     string           `json:"name"`
	ImageURL string           `json:"url"`
	Hotkeys  *[]hotkey.Hotkey `json:"hotkeys"`
}

type Repository interface {
	GetAll() (*[]Program, error)
	Get(id uint) (*Program, error)
	GetHotkeys(id uint) (*[]hotkey.Hotkey, error)
}
