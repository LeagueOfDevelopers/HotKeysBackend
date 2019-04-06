package getting

import (
	"HotKeysBackend/converter"
	"HotKeysBackend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleGetPrograms(context *gin.Context, getter Service) {
	programs, err := getter.GetPrograms()
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorGetProgram)
	}

	programsResponse := converter.ConvertProgramsToResponse(programs)
	context.JSON(http.StatusOK, gin.H{
		"programs": programsResponse,
	})
}

func HandleGetProgram(context *gin.Context, getter Service) {
	paramProgramId, err := strconv.ParseUint(context.Param(utils.ProgramId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongProgramIdFormat)
	}

	programId := uint(paramProgramId)
	currentProgram, err := getter.GetProgram(uint(programId))
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorGetProgram)
	}

	programResponse := converter.ConvertProgramToResponse(currentProgram)
	context.JSON(http.StatusOK, gin.H{
		"program": programResponse,
	})
}

func HandleGetHotkeys(context *gin.Context, getter Service) {
	paramProgramId, err := strconv.ParseUint(context.Param(utils.ProgramId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongProgramIdFormat)
	}

	programId := uint(paramProgramId)
	hotkeys, err := getter.GetHotkeysForProgram(programId)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorGetHotkeys)
	}

	hotkeysResponse := converter.ConvertHotkeysToResponse(hotkeys)
	context.JSON(http.StatusOK, gin.H{
		"hotkeys": hotkeysResponse,
	})
}
