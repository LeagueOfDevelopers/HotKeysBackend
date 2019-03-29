package getting

import (
	"HotKeysBackend/constant"
	"HotKeysBackend/converter"
	"HotKeysBackend/hotkey"
	"HotKeysBackend/program"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ErrorNoProgramIdFound = errors.New("no param for program id")

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
	queryParams := context.Request.URL.Query()

	if len(queryParams[constant.ProgramId]) == 0 {
		sendError(context, ErrorNoProgramIdFound)
		panic(program.ErrorProgramIdFormat)
	}

	paramProgramId, err := strconv.ParseUint(queryParams[constant.ProgramId][0], 10, 64)
	if err != nil {
		sendError(context, err)
		panic(program.ErrorProgramIdFormat)
	}

	programId := uint(paramProgramId)
	currentProgram, err := getter.GetProgram(programId)
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
	queryParams := context.Request.URL.Query()
	paramProgramId, err := strconv.ParseUint(queryParams[constant.ProgramId][0], 10, 64)
	if err != nil {
		sendError(context, err)
		panic(program.ErrorProgramIdFormat)
	}

	programId := uint(paramProgramId)
	hotkeys, err := getter.GetHotkeysForProgram(programId)
	if err != nil {
		sendError(context, err)
		panic(hotkey.ErrorHotkeysNotFound)
	}

	hotkeysResponse := converter.ConvertHotkeysToResponse(hotkeys)
	context.JSON(http.StatusOK, gin.H{
		"hotkeys": hotkeysResponse,
	})
}

func sendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"error": err.Error(),
	})
}
