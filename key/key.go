package key

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Key struct {
	gorm.Model
	Id   uuid.UUID `json:"id"`
	Code int       `json:"code"`
	Name string    `json:"name"`
}

var ErrorKeyNotFound = errors.New("key not found")

type Repository interface {
	GetAll() ([]*Key, error)
	Get(int) (*Key, error)
}
