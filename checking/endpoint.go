package checking

import (
	"HotKeysBackend/hotkey"
	"HotKeysBackend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleCheckHotkey(context *gin.Context, checker Service) {
	programId := getProgramId(context)
	hotkeyId := getHotkeyId(context)
	combination := make([]uint, hotkey.MAX_LENGTH)

	if context.Request.Body == nil {
		utils.SendError(context, nil)
		panic(utils.ErrorWrongBodyFormat)
	}

	buf := make([]byte, 1024)
	num, err := context.Request.Body.Read(buf)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongBodyFormat)
	}

	reqBody := string(buf[0:num])
	context.JSON(http.StatusOK, reqBody)

	right, err := checker.Check(programId, hotkeyId, combination)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorChecking)
	}

	if right {
		context.JSON(http.StatusOK, gin.H{"answer": true})
	} else {
		context.JSON(http.StatusOK, gin.H{"answer": false})
	}
}

func getProgramId(context *gin.Context) uint {
	paramProgramId, err := strconv.ParseUint(context.Param(utils.ProgramId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongProgramIdFormat)
	}
	return uint(paramProgramId)
}

func getHotkeyId(context *gin.Context) uint {
	paramHotkeyId, err := strconv.ParseUint(context.Param(utils.HotkeyId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongHotkeyIdFormat)
	}
	return uint(paramHotkeyId)
}
