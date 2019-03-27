package hotkey

import (
	"HotKeysBackend/key"
	"github.com/jinzhu/gorm"
)

type Hotkey struct {
	gorm.Model
	Description string     `json:"desc"`
	Combination []*key.Key `json:"combination"`
}
