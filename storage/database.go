package storage

import (
	"HotKeysBackend/key"
	"HotKeysBackend/program"
	"github.com/jinzhu/gorm"
)

const (
	// TODO find database
	databasePath = ""
)

func GetProgramDatabaseRepository() program.Repository {
	// TODO
	//database := getDatabase(databasePath)
	return nil
}

func GetKeyDatabaseRepository() key.Repository {
	// TODO
	//database := getDatabase(databasePath)
	return nil
}

func getDatabase(databasePath string) *gorm.DB {
	//db, err := gorm.Open(databasePath, "")
	//if err != nil {
	//	panic("failed to connect database")
	//}
	return nil
}
