package getting

import (
	"HotKeysBackend/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetPrograms(context *gin.Context, getter Service) {
	programs, err := getter.GetPrograms()
	if err != nil {
		sendError(context, err)
		panic("error while getting programs")
	}

	programsResponse := converter.ConvertProgramsToResponse(programs)
	context.JSON(http.StatusOK, gin.H{
		"programs": programsResponse,
	})
}

func HandleGetProgram(context *gin.Context, getter Service) {
	programName := context.Param("program")
	program, err := getter.GetProgram(programName)
	if err != nil {
		sendError(context, err)
		panic("error while getting program " + programName)
	}

	programResponse := converter.ConvertProgramToResponse(program)
	context.JSON(http.StatusOK, gin.H{
		"program": programResponse,
	})
}

func sendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"cause": err,
	})
}
