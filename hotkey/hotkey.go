package hotkey

import (
	"HotKeysBackend/key"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Hotkey struct {
	gorm.Model
	Id          uuid.UUID  `json:"id"`
	Description string     `json:"desc"`
	Combination []*key.Key `json:"combination"`
}
