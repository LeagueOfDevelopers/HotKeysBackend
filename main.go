package main

import (
	"HotKeysBackend/checking"
	"HotKeysBackend/constant"
	"HotKeysBackend/getting"
	"HotKeysBackend/storage"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	handleGetPrograms = "/programs"
	handleGetProgram  = fmt.Sprintf("/program/:%s", constant.ProgramName)
	handleGetHotkeys  = fmt.Sprintf("/program/:%s/hotkeys", constant.ProgramName)
	handleCheckHotkey = fmt.Sprintf("/program/:%s/check", constant.ProgramName)
)

func main() {
	// TODO uncomment when database appears
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
