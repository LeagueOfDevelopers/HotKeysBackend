package main

import (
	"HotKeysBackend/checking"
	"HotKeysBackend/getting"
	"HotKeysBackend/storage"
	"HotKeysBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	handleGetPrograms = "/programs/"
	handleGetProgram  = fmt.Sprintf("/programs/:%s", utils.ProgramId)
	handleGetHotkeys  = fmt.Sprintf("/programs/:%s/hotkeys", utils.ProgramId)
	handleCheckHotkey = fmt.Sprintf("/programs/:%s/hotkeys/:%s/check", utils.ProgramId, utils.HotkeyId)
)

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//defer db.Close()
	//
	//db.Create(&program.Program{})
	//programStorage := storage.GetProgramDatabaseRepository()
	//keyStorage := storage.GetKeyDatabaseRepository()

	programStorage := storage.GetProgramInMemoryRepository()
	keyStorage := storage.GetKeyInMemoryRepository()
	getter := getting.CreateService(programStorage, keyStorage)
	checker := checking.CreateService(programStorage, keyStorage)

	router := gin.Default()

	router.GET(handleGetPrograms, func(context *gin.Context) {
		getting.HandleGetPrograms(context, getter)
	})
	router.GET(handleGetProgram, func(context *gin.Context) {
		getting.HandleGetProgram(context, getter)
	})
	router.GET(handleGetHotkeys, func(context *gin.Context) {
		getting.HandleGetHotkeys(context, getter)
	})
	router.POST(handleCheckHotkey, func(context *gin.Context) {
		checking.HandleCheckHotkey(context, checker)
	})

	_ = router.Run()
}
