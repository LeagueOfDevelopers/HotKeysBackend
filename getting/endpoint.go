package getting

import (
	"HotKeysBackend/constant"
	"HotKeysBackend/converter"
	"HotKeysBackend/program"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetPrograms(context *gin.Context, getter Service) {
	programs, err := getter.GetPrograms()
	if err != nil {
		sendError(context, err)
		panic(program.ErrorGetProgram)
	}

	programsResponse := converter.ConvertProgramsToResponse(programs)
	context.JSON(http.StatusOK, gin.H{
		"programs": programsResponse,
	})
}

func HandleGetProgram(context *gin.Context, getter Service) {
	programName := context.Param(constant.ProgramName)
	currentProgram, err := getter.GetProgram(programName)
	if err != nil {
		sendError(context, err)
		panic(program.ErrorGetProgram)
	}

	programResponse := converter.ConvertProgramToResponse(currentProgram)
	context.JSON(http.StatusOK, gin.H{
		"program": programResponse,
	})
}

func HandleGetHotkeys(context *gin.Context, getter Service) {
	programName := context.Param(constant.ProgramName)
	hotkeys, err := getter.GetHotkeysForProgram(programName)
	if err != nil {
		sendError(context, err)
		panic(program.ErrorGetProgram)
	}

	hotkeysResponse := converter.ConvertHotkeysToResponse(hotkeys)
	context.JSON(http.StatusOK, gin.H{
		"hotkeys": hotkeysResponse,
	})
}

func sendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"cause": err,
	})
}
