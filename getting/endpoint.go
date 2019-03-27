package getting

import (
	"HotKeysBackend/converter"
	"HotKeysBackend/program"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ProgramName = "program"

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
	programName := context.Param(ProgramName)
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
	programName := context.Param(ProgramName)
	hotkeys, err := getter.GetHotkeysForProgram(programName)
	if err != nil {
		sendError(context, err)
		panic(program.ErrorGetProgram)
	}

	context.JSON(http.StatusOK, gin.H{
		"hotkeys": hotkeys,
	})
}

func sendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"cause": err,
	})
}
