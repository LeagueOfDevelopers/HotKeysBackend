package key

import (
	"github.com/jinzhu/gorm"
)

type Key struct {
	gorm.Model
	Code int    `json:"code"`
	Name string `json:"name"`
}

type Repository interface {
	GetAll() (*[]Key, error)
	Get(int) (*Key, error)
}
