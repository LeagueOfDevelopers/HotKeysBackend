package key

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Key struct {
	gorm.Model
	Code int    `json:"code"`
	Name string `json:"name"`
}

var ErrorKeyNotFound = errors.New("key not found")

type Repository interface {
	GetAll() ([]*Key, error)
	Get(int) (*Key, error)
}
