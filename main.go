package main

import (
	"HotKeysBackend/getting"
	"HotKeysBackend/storage"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	handleGetPrograms = "/programs"
	handleGetProgram  = fmt.Sprintf("/program/:%s", getting.ProgramName)
	handleGetHotkeys  = fmt.Sprintf("/program/:%s/hotkeys", getting.ProgramName)
)

func main() {
	// TODO uncomment when database appears
	//programStorage := storage.GetProgramDatabaseRepository()
	//keyStorage := storage.GetKeyDatabaseRepository()

	programStorage := storage.GetProgramInMemoryRepository()
	keyStorage := storage.GetKeyInMemoryRepository()
	getter := getting.NewService(programStorage, keyStorage)

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

	_ = router.Run()
}
