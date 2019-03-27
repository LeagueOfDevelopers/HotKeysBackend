package main

import (
	"HotKeysBackend/getting"
	"HotKeysBackend/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO uncomment when database appears
	//programStorage := storage.GetProgramDatabaseRepository()
	//keyStorage := storage.GetKeyDatabaseRepository()

	programStorage := storage.GetProgramInMemoryRepository()
	keyStorage := storage.GetKeyInMemoryRepository()
	getter := getting.NewService(programStorage, keyStorage)

	router := gin.Default()

	router.GET("/programs", func(context *gin.Context) {
		getting.HandleGetPrograms(context, getter)
	})
	router.GET("/program/:program", func(context *gin.Context) {
		getting.HandleGetProgram(context, getter)
	})

	_ = router.Run()
}
