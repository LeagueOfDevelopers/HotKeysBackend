package hotkey

import (
	"HotKeysBackend/key"
	"errors"
	"github.com/jinzhu/gorm"
)

type Hotkey struct {
	gorm.Model
	Description string     `json:"desc"`
	Combination *[]key.Key `json:"combination"`
}

var ErrorHotkeysNotFound = errors.New("hotkeys not found")
