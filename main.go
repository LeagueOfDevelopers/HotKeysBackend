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
	handleGetProgram  = fmt.Sprintf("/program/:%s", utils.ProgramId)
	handleGetHotkeys  = fmt.Sprintf("/program/:%s/hotkeys", utils.ProgramId)
	handleCheckHotkey = fmt.Sprintf("/program/:%s/hotkeys/:%s/check", utils.ProgramId, utils.HotkeyId)
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
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"POST", "GET", "PUT"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 4 * time.Hour,
	//}))

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
